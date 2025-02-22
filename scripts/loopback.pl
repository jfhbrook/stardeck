#!/usr/bin/env perl

use v5.40;
use feature qw(switch);
use warnings;

use Getopt::Long;

package main;

my $HELP = 'usage: loopback.pl [OPTIONS] ACTION

ACTIONS:
  enable   Enable loopback
  disable  Disable loopback

OPTIONS:
  --latency MSEC  Latency, in milliseconds
';

my $help    = '';
my $latency = 1;

GetOptions(
    "help"      => \$help,
    "latency=i" => \$latency,
) or die $HELP;

my $cmd;

if ($#ARGV) {
    die 'Must specify action.';
}
elsif ( $ARGV[0] eq 'enable' ) {
    $cmd = "pactl load-module module-loopback --latency_msec=$latency";
}
elsif ( $ARGV[0] eq 'disable' ) {
    $cmd = 'pactl unload-module module-loopback';
}
else {
    die "Unknown action $ARGV[0]";
}

my $status = system($cmd);

exit $status;
