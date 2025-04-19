package Stardeck::Window;

use 5.038;
use threads;
use threads::shared;
use warnings;

use Carp 'croak';
use IPC::Run 'run';
use Storable 'dclone';
use String::ShellQuote 'shell_quote';
use Time::HiRes 'usleep';

use constant WINDOW_POLL_INTERVAL => 200 * 10e3;

sub kdotool {
    my @command = @{ dclone( \@_ ) };
    unshift( @command, 'kdotool' );

    my $quoted = shell_quote @command;

    run( \@command, my $in, my $out, my $err )
      or croak "$quoted: $?";

    my $res = <$out>;

    chomp $res;

    while ( my $line = <$err> ) {
        print $line;
    }

    return $res;
}

sub get_window {
    return kdotool('getactivewindow');
}

sub get_window_name {
    my $window = get_window();
    return kdotool( 'getwindowname', $window );
}

my sub is_running {
    my $command = $_->dequeue_nb();
    return $command->{'type'} eq 'Stop';
}

sub window_worker {
    my $command_queue = shift;
    my $event_queue   = shift;

    my $current = '';
    my $running = is_running($command_queue);

    while ($running) {
        my $next = get_window_name();

        if ( $next ne $current ) {
            my %event = (
                type => 'ActiveWindow',
                name => "${next}"
            );

            $event_queue->enqueue( \%event );
        }
        $current = $next;

        $running = is_running($command_queue);

        if ($running) {
            usleep(WINDOW_POLL_INTERVAL);
        }

        $running = is_running($command_queue);
    }

    return;
}

1;
