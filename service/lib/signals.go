package lib

import (
	"github.com/godbus/dbus/v5"
)

// https://github.com/godbus/dbus/blob/master/_examples/signal.go
// https://pkg.go.dev/github.com/godbus/dbus/v5#Signal
func SignalListener(bus *dbus.Conn) {
	// Listens to all relevant signals from both crystalfontz and stardeck

}

/*
my sub plusdeck_state_worker {
    my @command = (
        $python_bin, '-u', '-m', 'plusdeck.dbus.client', '--output', 'text',
        'subscribe'
    );

    &listen_process(
        'plusdeck_state',
        \@command,
        sub {
            my %event = (
                type  => "PlusdeckState",
                state => $_
            );
            \%event;
        }
    );

    return;
}
*/

/*
my sub crystalfontz_reports_worker {
    my @command = (
        $python_bin, '-u', '-m', 'crystalfontz.dbus.client', '--output',
        'json',      'listen'
    );

    # {"type": "KeyActivityReport", "activity": "KEY_UP_PRESS"}
    &listen_process( 'crystalfontz_reports', \@command,
        sub { decode_json $_; } );

    return;
}
*/
