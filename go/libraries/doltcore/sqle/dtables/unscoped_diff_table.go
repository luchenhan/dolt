// Copyright 2022 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dtables

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/dolthub/dolt/go/libraries/doltcore/diff"
	"github.com/dolthub/dolt/go/libraries/doltcore/doltdb"
	"github.com/dolthub/dolt/go/libraries/doltcore/env"
	"github.com/dolthub/dolt/go/libraries/doltcore/env/actions/commitwalk"
	"github.com/dolthub/dolt/go/libraries/doltcore/sqle/dsess"
	"github.com/dolthub/dolt/go/store/datas"
	"github.com/dolthub/dolt/go/store/hash"

	"github.com/dolthub/go-mysql-server/sql"
)

var workingSetPartitionKey = []byte("workingset")
var commitHistoryPartitionKey = []byte("commithistory")

// UnscopedDiffTable is a sql.Table implementation of a system table that shows which tables have
// changed in each commit, across all branches.
type UnscopedDiffTable struct {
	ddb  *doltdb.DoltDB
	head *doltdb.Commit
}

// tableChange is an internal data structure used to hold the results of processing
// a diff.TableDelta structure into the output data for this system table.
type tableChange struct {
	tableName    string
	dataChange   bool
	schemaChange bool
}

// NewUnscopedDiffTable creates an UnscopedDiffTable
func NewUnscopedDiffTable(_ *sql.Context, ddb *doltdb.DoltDB, head *doltdb.Commit) sql.Table {
	return &UnscopedDiffTable{ddb: ddb, head: head}
}

// Name is a sql.Table interface function which returns the name of the table which is defined by the constant
// LogTableName
func (dt *UnscopedDiffTable) Name() string {
	return doltdb.DiffTableName
}

// String is a sql.Table interface function which returns the name of the table which is defined by the constant
// DiffTableName
func (dt *UnscopedDiffTable) String() string {
	return doltdb.DiffTableName
}

// Schema is a sql.Table interface function that returns the sql.Schema for this system table.
func (dt *UnscopedDiffTable) Schema() sql.Schema {
	return []*sql.Column{
		{Name: "commit_hash", Type: sql.Text, Source: doltdb.DiffTableName, PrimaryKey: true},
		{Name: "table_name", Type: sql.Text, Source: doltdb.DiffTableName, PrimaryKey: true},
		{Name: "committer", Type: sql.Text, Source: doltdb.DiffTableName, PrimaryKey: false},
		{Name: "email", Type: sql.Text, Source: doltdb.DiffTableName, PrimaryKey: false},
		{Name: "date", Type: sql.Datetime, Source: doltdb.DiffTableName, PrimaryKey: false},
		{Name: "message", Type: sql.Text, Source: doltdb.DiffTableName, PrimaryKey: false},
		{Name: "data_change", Type: sql.Boolean, Source: doltdb.DiffTableName, PrimaryKey: false},
		{Name: "schema_change", Type: sql.Boolean, Source: doltdb.DiffTableName, PrimaryKey: false},
	}
}

// Partitions is a sql.Table interface function that returns a partition of the data. Returns one
// partition for working set changes and one partition for all commit history.
func (dt *UnscopedDiffTable) Partitions(ctx *sql.Context) (sql.PartitionIter, error) {
	return NewSliceOfPartitionsItr([]sql.Partition{
		newDoltDiffPartition(workingSetPartitionKey),
		newDoltDiffPartition(commitHistoryPartitionKey),
	}), nil
}

// PartitionRows is a sql.Table interface function that gets a row iterator for a partition.
func (dt *UnscopedDiffTable) PartitionRows(ctx *sql.Context, partition sql.Partition) (sql.RowIter, error) {
	if bytes.Equal(partition.Key(), workingSetPartitionKey) {
		return dt.newWorkingSetRowItr(ctx)
	} else if bytes.Equal(partition.Key(), commitHistoryPartitionKey) {
		return dt.newCommitHistoryRowItr(ctx)
	} else {
		return nil, fmt.Errorf("unexpected partition: %v", partition)
	}
}

func (dt *UnscopedDiffTable) newWorkingSetRowItr(ctx *sql.Context) (sql.RowIter, error) {
	sess := dsess.DSessFromSess(ctx.Session)
	rootsProvider := dsess.NewSessionStateAdapter(sess.Session, ctx.GetCurrentDatabase(), map[string]env.Remote{}, map[string]env.BranchConfig{}, map[string]env.Remote{})

	roots, err := rootsProvider.GetRoots(ctx)
	if err != nil {
		return nil, err
	}

	staged, unstaged, err := diff.GetStagedUnstagedTableDeltas(ctx, roots)
	if err != nil {
		return nil, err
	}

	return &doltDiffWorkingSetRowItr{
		tableDeltas: append(staged, unstaged[0:]...),
	}, nil
}

var _ sql.RowIter = &doltDiffWorkingSetRowItr{}

type doltDiffWorkingSetRowItr struct {
	tableDeltas []diff.TableDelta
	idx         int
}

func (d *doltDiffWorkingSetRowItr) Next(ctx *sql.Context) (sql.Row, error) {
	if d.idx >= len(d.tableDeltas) {
		return nil, io.EOF
	}

	tableDelta := d.tableDeltas[d.idx]
	d.idx++

	dataChange := false
	schemaChange := false

	// TODO: is this the right logic? double check this?
	// TODO: Can we abstract and simplify any of this?
	if tableDelta.IsDrop() {
		rowData, err := tableDelta.FromTable.GetRowData(ctx)
		if err != nil {
			return nil, err
		}

		dataChange = rowData.Count() > 0
		schemaChange = true
	} else if tableDelta.IsAdd() {
		rowData, err := tableDelta.ToTable.GetRowData(ctx)
		if err != nil {
			return nil, err
		}

		dataChange = rowData.Count() > 0
		schemaChange = true
	} else {
		hashChanged, err := tableDelta.HasHashChanged()
		if err != nil {
			return nil, err
		}
		dataChange = hashChanged

		schemaChanged, err := tableDelta.HasSchemaChanged(ctx)
		if err != nil {
			return nil, err
		}
		schemaChange = schemaChanged
	}

	sqlRow := sql.NewRow(
		"WORKING",         // commit_hash
		tableDelta.ToName, // table_name
		"NULL",            // committer
		"NULL",            // email
		"NULL",            // date
		"NULL",            // message
		dataChange,        // data_change
		schemaChange,      // schema_change
	)

	return sqlRow, nil
}

func (d *doltDiffWorkingSetRowItr) Close(c *sql.Context) error {
	return nil
}

var _ sql.Partition = &doltDiffPartition{}

type doltDiffPartition struct {
	key []byte
}

func newDoltDiffPartition(key []byte) *doltDiffPartition {
	return &doltDiffPartition{
		key: key,
	}
}

func (d doltDiffPartition) Key() []byte {
	return d.key
}

// doltDiffCommitHistoryRowItr is a sql.RowItr implementation which iterates over each commit as if it's a row in the table.
type doltDiffCommitHistoryRowItr struct {
	ctx             *sql.Context
	ddb             *doltdb.DoltDB
	child           doltdb.CommitItr
	meta            *datas.CommitMeta
	hash            hash.Hash
	tableChanges    []tableChange
	tableChangesIdx int
}

// newCommitHistoryRowItr creates a doltDiffCommitHistoryRowItr from the current environment.
func (dt *UnscopedDiffTable) newCommitHistoryRowItr(ctx *sql.Context) (*doltDiffCommitHistoryRowItr, error) {
	hash, err := dt.head.HashOf()
	if err != nil {
		return nil, err
	}
	child, err := commitwalk.GetTopologicalOrderIterator(ctx, dt.ddb, hash)

	return &doltDiffCommitHistoryRowItr{
		ctx:             ctx,
		ddb:             dt.ddb,
		child:           child,
		tableChangesIdx: -1,
	}, nil
}

// incrementIndexes increments the table changes index, and if it's the end of the table changes array, moves
// to the next commit, and resets the table changes index so that it can be populated when Next() is called.
func (itr *doltDiffCommitHistoryRowItr) incrementIndexes() {
	itr.tableChangesIdx++
	if itr.tableChangesIdx >= len(itr.tableChanges) {
		itr.tableChangesIdx = -1
		itr.tableChanges = nil
	}
}

// Next retrieves the next row. It will return io.EOF if it's the last row.
// After retrieving the last row, Close will be automatically closed.
func (itr *doltDiffCommitHistoryRowItr) Next(ctx *sql.Context) (sql.Row, error) {
	defer itr.incrementIndexes()

	for itr.tableChanges == nil {
		err := itr.loadTableChanges(ctx)
		if err != nil {
			return nil, err
		}
	}

	tableChange := itr.tableChanges[itr.tableChangesIdx]
	meta := itr.meta
	h := itr.hash

	return sql.NewRow(
		h.String(),
		tableChange.tableName,
		meta.Name,
		meta.Email,
		meta.Time(),
		meta.Description,
		tableChange.dataChange,
		tableChange.schemaChange,
	), nil
}

// loadTableChanges loads the next commit's table changes and metadata
// into the iterator.
func (itr *doltDiffCommitHistoryRowItr) loadTableChanges(ctx context.Context) error {
	hash, commit, err := itr.child.Next(ctx)
	if err != nil {
		return err
	}

	tableChanges, err := itr.calculateTableChanges(ctx, commit)
	if err != nil {
		return err
	}

	itr.tableChanges = tableChanges
	itr.tableChangesIdx = 0
	if len(tableChanges) == 0 {
		return nil
	}

	meta, err := commit.GetCommitMeta(ctx)
	if err != nil {
		return err
	}
	itr.meta = meta
	itr.hash = hash
	return nil
}

// calculateTableChanges calculates the tables that changed in the specified commit, by comparing that
// commit with its immediate ancestor commit.
func (itr *doltDiffCommitHistoryRowItr) calculateTableChanges(ctx context.Context, commit *doltdb.Commit) ([]tableChange, error) {
	if len(commit.DatasParents()) == 0 {
		return nil, nil
	}

	toRootValue, err := commit.GetRootValue(ctx)
	if err != nil {
		return nil, err
	}

	parent, err := itr.ddb.ResolveParent(ctx, commit, 0)
	if err != nil {
		return nil, err
	}

	fromRootValue, err := parent.GetRootValue(ctx)
	if err != nil {
		return nil, err
	}

	deltas, err := diff.GetTableDeltas(ctx, fromRootValue, toRootValue)
	if err != nil {
		return nil, err
	}

	tableChanges := make([]tableChange, len(deltas))
	for i := 0; i < len(deltas); i++ {
		change, err := itr.processTableDelta(deltas[i])
		if err != nil {
			return nil, err
		}

		tableChanges[i] = *change
	}

	// Not all commits mutate tables (e.g. empty commits)
	if len(tableChanges) == 0 {
		return nil, nil
	}

	return tableChanges, nil
}

// processTableDelta processes the specified TableDelta to determine what kind of change it was (i.e. table drop,
// table rename, table create, or data update) and returns a tableChange struct representing the change.
func (itr *doltDiffCommitHistoryRowItr) processTableDelta(delta diff.TableDelta) (*tableChange, error) {
	// Dropping a table is always a schema change, and also a data change if the table contained data
	if itr.isTableDropChange(delta) {
		isEmpty, err := itr.isTableDataEmpty(delta.FromTable)
		if err != nil {
			return nil, err
		}

		return &tableChange{
			tableName:    delta.FromName,
			dataChange:   !isEmpty,
			schemaChange: true,
		}, nil
	}

	// Renaming a table is always a schema change, and also a data change if the table data differs
	if delta.IsRename() {
		dataChanged, err := delta.HasHashChanged()
		if err != nil {
			return nil, err
		}

		return &tableChange{
			tableName:    delta.ToName,
			dataChange:   dataChanged,
			schemaChange: true,
		}, nil
	}

	// Creating a table is always a schema change, and also a data change if data was inserted
	if itr.isTableCreateChange(delta) {
		isEmpty, err := itr.isTableDataEmpty(delta.ToTable)
		if err != nil {
			return nil, err
		}

		return &tableChange{
			tableName:    delta.ToName,
			dataChange:   !isEmpty,
			schemaChange: true,
		}, nil
	}

	dataChanged, err := delta.HasHashChanged()
	if err != nil {
		return nil, err
	}

	schemaChanged, err := delta.HasSchemaChanged(itr.ctx)
	if err != nil {
		return nil, err
	}

	return &tableChange{
		tableName:    delta.ToName,
		dataChange:   dataChanged,
		schemaChange: schemaChanged,
	}, nil
}

// Close closes the iterator.
func (itr *doltDiffCommitHistoryRowItr) Close(*sql.Context) error {
	return nil
}

// isTableDataEmpty return true if the table does not contain any data
func (itr *doltDiffCommitHistoryRowItr) isTableDataEmpty(table *doltdb.Table) (bool, error) {
	rowData, err := table.GetRowData(itr.ctx)
	if err != nil {
		return false, err
	}

	return rowData.Empty(), nil
}

// isRenameChange returns true if the specified TableDelta represents a table rename change.
func (itr *doltDiffCommitHistoryRowItr) isRenameChange(delta diff.TableDelta) bool {
	return delta.FromTable != nil &&
		delta.ToTable != nil &&
		delta.FromName != delta.ToName
}

// isTableDropChange return true if the specified TableDelta represents a table drop change.
func (itr *doltDiffCommitHistoryRowItr) isTableDropChange(delta diff.TableDelta) bool {
	return len(delta.FromName) > 0 && len(delta.ToName) == 0
}

// isTableCreateChange returns true if the specified TableDelta represents a table create change.
func (itr *doltDiffCommitHistoryRowItr) isTableCreateChange(delta diff.TableDelta) bool {
	return delta.FromTable == nil && delta.ToTable != nil
}
