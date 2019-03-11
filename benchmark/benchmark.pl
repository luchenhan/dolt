#!/usr/bin/perl -w

use strict;

use Data::Dumper;

# These are defaults and will al be able to be overridden with command line args.
use constant BENCHMARK_ROOT => '/var/tmp';
use constant VERBOSE => 0;

use constant TEST_FILE      => 'test.csv'; 
use constant TEST_INPUT_CSV => BENCHMARK_ROOT . '/' . TEST_FILE;

# Set up the environment
# TO DO: Figure out a portable way to get dolt in the path
$ENV{'PATH'} = $ENV{'PATH'} . ':~/go/bin/';
$ENV{'NOMS_VERSION_NEXT'} = 1;

my $changes = [
    {
        filehandle => *SMALL,
        file => BENCHMARK_ROOT . '/small-change.csv',
        pct  => 0.01,
    },
    {
        filehandle => *MEDIUM,
        file => BENCHMARK_ROOT . '/medium-change.csv',
        pct  => 0.1,
    },
    {
        filehandle => *LARGE,
        file => BENCHMARK_ROOT . '/large-change.csv',
        pct  => 0.5,
    },
];

# Define the benchmarks we will run.
my $benchmarks = {
    git => {
	root => BENCHMARK_ROOT . '/git-benchmark/',
	tests => [
	    {
		name => 'raw',
		command => 'git',
	    },
	    {
		name => 'init',
		command => 'git init',
	    },

	    {
		prep => [
		    'cp ' . TEST_INPUT_CSV . ' ' . BENCHMARK_ROOT . '/git-benchmark/',
		    ],
		name => 'add',
		command => 'git add ' . TEST_FILE,
	    },
            {
                name => 'commit',
                command => 'git commit -m "first test commit"',
            },
	    {
                prep => [
                    'cp ' . $changes->[0]{'file'} . ' ' . BENCHMARK_ROOT . '/git-benchmark/' . TEST_FILE,
                    ],
                name => 'small diff',
                command => 'git diff ' . TEST_FILE,
            },
            {
                prep => [
                    'cp ' . $changes->[1]{'file'} . ' ' . BENCHMARK_ROOT . '/git-benchmark/' . TEST_FILE,
                    ],
                name => 'medium diff',
                command => 'git diff ' . TEST_FILE,
            },
            {
                prep => [
                    'cp ' . $changes->[2]{'file'} . ' ' . BENCHMARK_ROOT . '/git-benchmark/' . TEST_FILE,
                    ],
                name => 'large diff',
                command => 'git diff ' . TEST_FILE,
            },
	],
    },
    dolt => {
	root => BENCHMARK_ROOT . '/dolt-benchmark/',
        tests => [
	    {
		name => 'raw',
		command => 'dolt',
	    },
	    {
		name => 'init',
		command => 'dolt init',
	    },
            {
		# Need to set up the schema here.
		prep => ['dolt table import -c --pk=id test ' . TEST_INPUT_CSV],
		name =>'add',
		command=> 'dolt add test',
	    },
            {
                name => 'commit',
                command => 'dolt commit -m "first test commit"',
            },
	    {
                prep => ['dolt table import -u test ' . $changes->[0]{'file'}],
                name => 'small diff',
                command=> 'dolt diff test',
            },
            {
                prep => ['dolt table import -u test ' . $changes->[1]{'file'}],
                name => 'medium diff',
                command=> 'dolt diff test',
            },
            {
                prep => ['dolt table import -u test ' . $changes->[2]{'file'}],
                name => 'large diff',
                command=> 'dolt diff test',
            },
	],
    },
};

# Define the schema and size of the test database
my $lines = 1000000;
my $schema = [
    {
	name    => 'id',
	type    => 'int',
	primary => 1,
	gen     => 'increment',
    },
    {
	name    => 'int1',
        type    => 'int',
        primary => 0,
	gen     => 'rand',
	size    => 10,
    },
    {
	name    => 'int2',
	type    => 'int',
        primary => 0,
        gen     => 'rand',
	size    => 100,
    },
    {
	name    => 'int3',
	type    => 'int',
        primary => 0,
        gen     => 'rand',
	size    => 1000,
    },
    {
	name    => 'int4',
	type    => 'int',
        primary => 0,
        gen     => 'rand',
	size    => 10000,
    },
    {
	name    => 'int5',
	type    => 'int',
        primary => 0,
        gen     => 'rand',
	size    => 100000,
    },
    {
	name    => 'string1',
	type    => 'string',
        primary => 0,
        gen     => 'rand',
	size    => 1,
    },
    {
	name    => 'string2',
        type    => 'string',
        primary => 0,
        gen     => 'rand',
        size    => 2,
    },
    {
	name    => 'string3',
        type    => 'string',
        primary => 0,
        gen     => 'rand',
        size    => 4,
    },
    {
	name    => 'string4',
        type    => 'string',
        primary => 0,
        gen     => 'rand',
        size    => 8,
    },
    {
	name    => 'string5',
        type    => 'string',
        primary => 0,
        gen     => 'rand',
        size    => 16,
    },
];

# Bootstrap the test
if ( -d BENCHMARK_ROOT ) { 
    chdir(BENCHMARK_ROOT);
} else {
    die "Could not run benchmarks in " . BENCHMARK_ROOT . 
	" because the directory does not exist.";
}

create_test_input_csvs(TEST_INPUT_CSV, $lines, $schema, $changes);

# Run the benchmarks
my %output;
foreach my $benchmark ( keys %{$benchmarks} ) {
    print "Executing $benchmark benchmark...\n" if VERBOSE;

    my $root = $benchmarks->{$benchmark}{'root'};
    if ( -d $root ) {
	run_command("rm -rf $root", VERBOSE);
	# die "$root must not exist to run benchmark\n";
    } else {
	mkdir($root);
	chdir($root);
	print "Changing directory to $root\n" if VERBOSE;
    }

    foreach my $test ( @{$benchmarks->{$benchmark}{'tests'}} ) {
	foreach my $prep ( @{$test->{'prep'}} ) {
	    run_command($prep, VERBOSE);
	}
	
	print "Running test: " . $test->{'name'} . "\n" if VERBOSE;

	my ($real, $user, $system) = time_command($test->{'command'}, VERBOSE);

	$output{$test->{'name'}}{$benchmark}{'real'}   = $real;
	$output{$test->{'name'}}{$benchmark}{'user'}   = $user;
	$output{$test->{'name'}}{$benchmark}{'system'} = $system;
    }

    run_command("rm -rf $root", VERBOSE);
}

# Cleanup
# unlink(TEST_INPUT_CSV);

print Dumper(\%output);

###################################################################################
#
# Functions
#
###################################################################################

sub time_command {
    my $command = shift;
    my $verbose = shift || 0; 

    print "Running:\n\t$command\n" if $verbose;

    # time outputs to STDERR so I'll trash STDOUT and grab STDERR from
    # STDOUT which `` writes to
    my $piped_command;
    if ( $verbose ) {
	$piped_command = "{ time $command ;} 2>&1";
    } else {
	$piped_command = "{ time $command ;} 2>&1 1>/dev/null";
    }

    my $output = `$piped_command`;
    # To Do: Some of these commands expect to exit 1. ie, git and dolt.
    # I need to build in an expect into the benchmark definition
    # if ($?) {
    #     die "Error running: $piped_command\n";
    # }

    $output =~ /real\s+(.+)\nuser\s+(.+)\nsys\s+(.+)\n/;

    print "Output:\n\t$output\n" if ( $verbose and $output );

    my $real   = convert_time_output_to_ms($1);
    my $user   = convert_time_output_to_ms($2);
    my $system = convert_time_output_to_ms($3);

    return ($real, $user, $system);
}

sub run_command {
    my $command = shift;
    my $verbose = shift || 0;

    print "Running:\n\t$command\n" if $verbose;
    my $output = `$command 2>&1`;
    print "Output:\n\t$output\n" if ( $verbose and $output );
    if ($?) {
	die "Error running: $command\n";
    }
}

sub convert_time_output_to_ms {
    my $time = shift;

    $time =~ /(\d+)m(\d+)\.(\d+)s/;

    my $minutes = $1;
    my $seconds = $2;
    my $ms      = $3;

    return $ms + ($seconds*1000) + ($minutes*60*1000);
}

sub create_test_input_csvs {
    my $csv     = shift;
    my $size    = shift;
    my $schema  = shift;
    my $changes = shift;

    my @all_filehandles;
    open(CSV, ">", $csv) or die "Could not open $csv: $!\n";
    push @all_filehandles, *CSV;

    foreach my $change ( @{$changes} ){
	open($change->{'filehandle'}, '>', $change->{'file'}) 
	    or die "Could not open ". $change->{'file'} . ": $!\n";
	push @all_filehandles, $change->{'filehandle'};
    }

    # Create header row and write it to all csvs
    my $first = 1;
    foreach my $column ( @{$schema} ) {
	write_to_files(',', @all_filehandles) unless $first; 
	write_to_files($column->{'name'}, @all_filehandles);
	$first = 0;
    }
    write_to_files("\n", @all_filehandles);;

    # Create mock data
    for ( my $i=0; $i < $size; $i++ ) {
	$first = 1;
	foreach my $column ( @{$schema} ) {
	    write_to_files(',', @all_filehandles) unless $first;
	    my $value = generate_value($column->{'type'}, 
				       $column->{'gen'}, 
				       $column->{'size'}, 
				       $i);
	    print CSV $value;
	    
            # Change the value if we roll under the probability
	    foreach my $change ( @{$changes} ) {
		if ( rand() < $change->{'pct'} ) {
		    $value = generate_value($column->{'type'},
					    $column->{'gen'},
					    $column->{'size'},
					    $i);
		} 
		my $fh = $change->{'filehandle'};
		print $fh $value;
	    }
	    $first = 0;
	}
	write_to_files("\n", @all_filehandles);
    }

    foreach my $fh (@all_filehandles) {
	close $fh;
    }
}

sub generate_value {
    my $type = shift;
    my $gen  = shift;
    my $size = shift;
    my $i    = shift; # Used for increment

    if ( $type eq 'int' ) {
	return $i if ( $gen eq 'increment' );
	if ( $gen eq 'rand' ) {
	    return int(rand($size+1));
	} else {
	    die "Do not understand generator: $gen\n";
	}
    } elsif ( $type eq 'string' ) {
	if ( $gen eq 'rand' ) {
	    return rndStr($size, 'a'..'z', 0..9);
	} else {
            die"Do not understand generator: $gen\n";
	}
    } else {
	die "Do not understand type: $type\n";
    }
}

sub write_to_files {
    my $string = shift;
    my @filehandles = @_;

    foreach my $filehandle ( @filehandles ) {
	print $filehandle $string;
    }
}

# Perl wizardry. Do not question.
sub rndStr { 
    join('', @_[ map{ rand @_ } 1 .. shift ]); 
}
