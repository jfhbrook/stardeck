package lib

/*
my sub listen_process {
    my $process_name = shift;
    my $command      = shift;
    my $map_sub      = shift;

    # TODO: Use open3 and log errors
    my $pid = open2( my $out, my $in, @$command );

    my %pid_event = (
        type => 'StartProcess',
        name => $process_name,
        pid  => $pid
    );

    $main_commands->enqueue( \%pid_event );

    while ($running) {
        my $line = <$out>;
        chomp $line;

        my $event = eval { $map_sub->($line); };

        if ($event) {
            $main_commands->enqueue($event);
        }
        else {
            # TODO: Log
            say $line;
        }
    }

    if ( waitpid( $pid, 0 ) > 0 ) {
        my $code = $?;

        my %exit_event = (
            type => 'ExitProcess',
            name => $process_name,
            code => $code
        );

        $main_commands->enqueue( \%exit_event );
    }

    return;
}

*/
