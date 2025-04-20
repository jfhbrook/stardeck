#!perl
use 5.038;
use strict;
use warnings;
use Test::More;

plan tests => 3;

BEGIN {
    use_ok('Stardeck')          || print "Bail out!\n";
    use_ok('Stardeck::Process') || print "Bail out!\n";
    use_ok('Stardeck::Window')  || print "Bail out!\n";
}

diag("Testing Stardeck $Stardeck::VERSION, Perl $], $^X");
