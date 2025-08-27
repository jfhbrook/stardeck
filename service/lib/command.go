package lib

/*
my %commands = (
    StartProcess => sub {
        my $name = $_->{'name'};
        my $pid  = $_->{'pid'};

        $pids{$name} = $pid;
    },
    ExitProcess => sub {

        # Clean up the PID
        my $name = $_->{'name'};
        delete $pids{$name};

        # A process exiting is unexpected, so we need to error out. Still,
        # we can killl the other processes gracefully.
        &graceful_exit(1);

    },
    ActiveWindow => sub {
        print Dumper($_);
    },
    PlusdeckState => sub {
        print Dumper($_);
    },
    KeyActivityReport => sub {
        print Dumper($_);
    },
    TemperatureReport => sub {
        print Dumper($_);
    },
    Notification => sub {
        print Dumper($_);
    }
);

my sub command_worker {
    while ( my $event = $main_commands->dequeue() ) {
        my $type    = $event->{'type'};
        my $handler = $commands{$type};

        if ($handler) {
            $handler->($event);
        }
        else {
            print Dumper($event);
        }
    }

    say "Unreachable.";
    &hard_exit(1);

    return;
}
*/
