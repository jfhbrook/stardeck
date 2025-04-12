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
  --latency MSEC   Loopback latency, in milliseconds
  --source SOURCE  The loopback source ID
  --volume VOLuME  The volume of the loopback source, in percent
';

my $help    = '';
my $latency = 1;
my $source  = 'alsa_input.pci-0000_00_1f.3.analog-stereo';
my $volume  = 10;

GetOptions(
    "help"      => \$help,
    "latency=i" => \$latency,
    "source=s"  => \$source,
    "volume=i"  => \$volume,
) or die $HELP;

my $load_cmd;
my $volume_cmd = "";

if ($#ARGV) {
    die 'Must specify action.';
}
elsif ( $ARGV[0] eq 'enable' ) {
    $load_cmd = "pactl load-module module-loopback --latency_msec=$latency";
    $volume_cmd =
"pactl set-source-volume alsa_input.pci-0000_00_1f.3.analog-stereo $volume%";
}
elsif ( $ARGV[0] eq 'disable' ) {
    $load_cmd = 'pactl unload-module module-loopback';
}
else {
    die "Unknown action $ARGV[0]";
}

my $status = system($load_cmd);

if ($status) {
    exit $status;
}

if ($volume_cmd) {
    $status = system($volume_cmd);
    exit $status;
}
