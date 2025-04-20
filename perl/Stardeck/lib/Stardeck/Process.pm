package Stardeck::Process;

# ABSTRACT: Process helpers for Stardeck

use 5.038;
use threads;
use threads::shared;
use warnings;

use Carp 'croak';
use IPC::Run 'run';
use Storable 'dclone';
use String::ShellQuote 'shell_quote';

our $VERSION = '0.01';

sub run_command {
    my @command = @{ dclone( \@_ ) };

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

1;
