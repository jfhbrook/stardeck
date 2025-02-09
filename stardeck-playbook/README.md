# stardeck-playbook

This component runs ansible playbooks. It will be written in perl.

## TODOs

- Set up a perl environment, locally and on the stardeck
  - [carton](https://metacpan.org/pod/Carton)
- Parse CLI options/flags
  - [Getopt::Long](https://perldoc.perl.org/Getopt::Long)
    - Stdlib, a little like argparse
  - [Getopt::Long::Descriptive](https://metacpan.org/pod/Getopt::Long::Descriptive)
    - Flavored like `Getopt::Long`, but with extra niceties. Maybe a good middle ground.
  - [App::Cmd](https://metacpan.org/pod/App::Cmd)
    - A little magical, but not in an offensive way. Looks fun?
  - [MooseX::App](https://metacpan.org/pod/MooseX::App)
    - [MooseX appears to be "CLOS but for Perl", yikes](https://www.perl.org/about/whitepapers/perl-object-oriented.html)
- Create/modify default/example yaml configs
- Load yaml config
- How to handle become pass?
  - Can I just run the playbooks as sudo?
  - Expect? <https://metacpan.org/release/RGIERSIG/Expect-1.15/view/Expect.pod>
  - Can developer run stardeck-playbook in an unprivileged mode? Just for their own environment?
- How to run commands in parallel with perl?
  - <https://stackoverflow.com/questions/54389215/run-multiple-jobs-within-perl-script-at-the-same-time>
- Init/generate default yaml config
  - Nice prompts?
  - Maybe not a feature for `stardeck-playbook`, per se? Better for `stardeckctl` or `stardeck-config`?
