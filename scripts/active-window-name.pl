#!/usr/bin/env perl

use v5.40;
use warnings;

package main;

my sub active_window_name {
    my $active_window = `kdotool getactivewindow`;
    `kdotool getwindowname $active_window`;
}

my $name = &active_window_name;

say $name;
