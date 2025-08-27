package lib

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
