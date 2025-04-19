# Stardeck Perl Code

Pardon our dust!

## develop phase

I have dev deps. I think I can configure those with:

```perl
on 'develop' => sub {
    requires ...
}
```

That may even allow for using the `cpanfile` to configure `dzil` based dependencies, instead of doing it auto.

## global tools

Right now, each project installs its own version of `dzil`. That's less than ideal.

I don't think there's a good way around it for `dzil`, since I use it for installing stuff into the carton environment. That said, here are some ideas...

I think that, if I set up perlbrew:

<https://perlbrew.pl/>

that I can then install those "global" tools with [cpanm](https://metacpan.org/pod/App::cpanminus):

```sh
cpanm Dist::Zilla
```

I may even be able to get it to install "global" dependencies with the `--cpanfile` flag.

I don't think perlbrew supports "gemsets". But this should be good enough.

## workspaces

Something I'm struggling with is workspaces. Neither Carton nor `dzil` support them, really. You can run `dzil install` in a project to install it in the current environment - that will let you fake it.
