package Stardeck;

use 5.040;
use threads;
use threads::shared;
use warnings;

use IPC::Run 'run';
use String::ShellQuote 'shell_quote';
use Time::HiRes 'usleep';

=head1 NAME

Stardeck - The great new Stardeck!

=head1 VERSION

Version 0.01

=cut

our $VERSION = '0.01';


=head1 SYNOPSIS

Quick summary of what the module does.

Perhaps a little code snippet.

    use Stardeck;

    my $foo = Stardeck->new();
    ...

=head1 EXPORT

A list of functions that can be exported.  You can delete this section
if you don't export anything, such as for a purely object-oriented module.

=head1 SUBROUTINES/METHODS

=head2 kdotool

=cut

sub kdotool {
    unshift(@_, 'kdotool');

    my $quoted = shell_quote @_;

    run( \@_, my $in, my $out, my $err )
      or die "$quoted: $?";

    my $res = <$out>;

    chomp $res;

    foreach (<$err>) {
        print $_;
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
    return kdotool('getwindowname', $window);
}

=head2 window_worker

=cut

my sub is_running {
    my $command = $_->dequeue_nb();
    return $command->{'type'} eq 'Stop';
}

sub window_worker {
    my $command_queue = shift;
    my $event_queue = shift;

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

1; # End of Stardeck
