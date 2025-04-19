# Stardeck Perl Code

Pardon our dust!

## Installing Local Dependencies

A problem I haven't solved is setting up local dependencies for a module. For instance, `Stardeck::Window` will depend on `Stardeck::Process`, but `Stardeck::Process` isn't avilable on CPAN.

It doesn't look like the [cpanfile](https://metacpan.org/dist/Module-CPANfile/view/lib/cpanfile.pod) format supports installing from local directory. Which is a shame. But `carton` will at least create/manage an environment.

You can install a local dzil module with `dzil install`. This includes all my authored modules, so that should work. But installing the *current* module into the carton environment causes breakage. So each module will need its own environment.
