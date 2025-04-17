package Stardeck;

use 5.040;
use threads;
use threads::shared;
use warnings;

use Croak 'croak';
use IPC::Run 'run';
use Storable 'clone';
use String::ShellQuote 'shell_quote';
use Time::HiRes 'usleep';

=head1 NAME

Stardeck - Interface with the Stardeck 1A Media Appliance

=head1 VERSION

Version 0.01

=cut

our $VERSION = '0.01';

=head1 SYNOPSIS

A module for interacting with the Stardeck 1A Media Appliance.

=head1 SUBROUTINES/METHODS

=head2 kdotool

=cut

sub kdotool {
    my @command = @{ clone( \@_ ) };
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

=head2 get_window

=cut

sub get_window {
    return kdotool('getactivewindow');
}

=head2 get_window_name

=cut

sub get_window_name {
    my $window = get_window();
    return kdotool( 'getwindowname', $window );
}

=head2 window_worker

=cut

my sub is_running {
    my $command = $_->dequeue_nb();
    return $command->{'type'} eq 'Stop';
}

sub window_worker {
    my $command_queue = shift;
    my $event_queue   = shift;

    my $running = is_running($command_queue);

    while ($running) {
        my $next = get_window_name();

        if ( $next ne $current ) {
            my %event = (
                type => 'ActiveWindow',
                name => "${next}"
            );

            $event_queoe->enqueue( \%event );
        }
        $current = $next;

        $running = is_running($command_queue);

        if ($running) {
            usleep( $active_window_poll_interval * 10e6 );
        }

        $running = is_running($command_queue);
    }

    return;
}

=head1 AUTHOR

Josh Holbrook, C<< <josh.holbrook at gmail.com> >>

=head1 BUGS

Please report any bugs or feature requests to C<bug-stardeck at rt.cpan.org>, or through
the web interface at L<https://rt.cpan.org/NoAuth/ReportBug.html?Queue=Stardeck>.  I will be notified, and then you'll
automatically be notified of progress on your bug as I make changes.




=head1 SUPPORT

You can find documentation for this module with the perldoc command.

    perldoc Stardeck


You can also look for information at:

=over 4

=item * RT: CPAN's request tracker (report bugs here)

L<https://rt.cpan.org/NoAuth/Bugs.html?Dist=Stardeck>

=item * CPAN Ratings

L<https://cpanratings.perl.org/d/Stardeck>

=item * Search CPAN

L<https://metacpan.org/release/Stardeck>

=back


=head1 ACKNOWLEDGEMENTS


=head1 LICENSE AND COPYRIGHT

This software is Copyright (c) 2025 by Josh Holbrook.

This is free software, licensed under:

  The Artistic License 2.0 (GPL Compatible)


=cut

1;    # End of Stardeck
