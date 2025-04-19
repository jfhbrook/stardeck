#!perl
use 5.038;
use strict;
use warnings;
use Test::More;

plan tests => 1;

BEGIN {
    use_ok('Stardeck::Window') || print "Bail out!\n";
}

diag("Testing Stardeck::Window $Stardeck::Window::VERSION, Perl $], $^X");
