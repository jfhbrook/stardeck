#!perl
use 5.038;
use strict;
use warnings;
use Test::More;

plan tests => 1;

BEGIN {
    use_ok('Stardeck::Process') || print "Bail out!\n";
}

diag("Testing Stardeck::Process $Stardeck::Process::VERSION, Perl $], $^X");
