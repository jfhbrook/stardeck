#!/usr/bin/env perl

use 5.040;
use threads;
use threads::shared;
use warnings;

use experimental 'class';

use Data::Dumper;
use IPC::Run 'run';
use Time::HiRes 'usleep';

# TODO: Use https://perldoc.perl.org/constant
my $active_window_poll_interval = 0.2;    # in seconds

my sub get_window {
    my @command = qw(kdotool getactivewindow);
    run( \@command, my $in, my $out, my $err )
      or die "kdotool getactivewindow: $?";

    my $window = <$out>;

    chomp $window;

    foreach (<$err>) {
        print $_;
    }

    return $window;
}

my sub get_window_name {
    my $window  = get_window();
    my @command = ( 'kdotool', 'getwindowname', $window );
    run( \@command, my $in, my $out, my $err )
      or die "kdotool getwindowname $window: $?";

    my $name = <$out>;

    chomp $name;

    foreach (<$err>) {
        print $_;
    }

    return $name;
}

my class Stardeck::ActiveWindowEmitter {
    field $running;
    field $interval;
    field $window;
    field $events : param;

    ADJUST {
        $running  = 1;
        $interval = 0.2;    # seconds
        $window   = '';
    }

    method stop {
        $self->running = 0;
    }

    method poll {
        my $next = get_window_name();
        my %event;
        if ( $next ne $self->window ) {
            %event = (
                type => 'ActiveWindow',
                name => "${next}"
            );

        }
        $self->window = $next;

        return \%event;
    }

    method run {
        while ( $self->running ) {
            my $event = $self->poll();

            if ($event) {
                $self->events->enqueue($event);
            }

            if ( $self->running ) {
                usleep( $self->interval * 10e6 );
            }
        }
    }
}
