package lib

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)

func Service() error {
	windowPollInterval := 0.2

	sessionConn, err := dbus.ConnectSessionBus()

	if err != nil {
		return err
	}

	systemConn, err := dbus.ConnectSystemBus()

	if err != nil {
		return err
	}

	events := make(chan *Event)

	go ListenToWindow(windowPollInterval, &events)
	go ListenToSignals(systemConn, &events)
	go ListenToNotifications(sessionConn, &events)

	for {
		event := <-events
		fmt.Println(event)
	}
}

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
