# stardeck-playbook

This tool runs Ansible playbooks to support the StarDeck 1A Media Appliance.

## Requirements

These playbooks are intended to be run on a PC running Fedora Linux, with a Plus Deck 2C PC Cassette Deck and a Crystalfontz CF533 LCD display. In other words, if you don't have my specific device, you will probably have a bad time.

## Usage

This tool is intended to be run with `npm exec` and `sudo`:

```sh
sudo npm exec --yes @jfhbrook/stardeck-playbook@latest
```

The tool supports a number of command line options:

```
USAGE: stardeck-playbook OPTIONS

OPTIONS:
  -h|--help              Show this help text and exit.
  --ansible-config FILE  The path to an ansible.cfg config file.
  --config-file FILE     The path to a stardeck.yml config file.
  --dry-run              Run ansible with --check and --diff.
  --feature FEATURE      Target a feature. May be specified more than once.
  --log-level LEVEL      Set the log level. Valid values are: debug, verbose, info, warn, error
  --no-update            Do not run software updates.
  --serial               Run playbooks in order.
  --use-version          Specify the version of stardeck-playbook to run.

ENVIRONMENT:
  ANSIBLE_CONFIG         A path to an ansible.cfg configuration file.
  STARDECK_CONFIG_FILE   A path to a stardeck.yml configuration file.
  STARDECK_CONFIG_HOME   A directory containing stardeck configuration files.

FEATURES:
  core          Core repositories and packages.
  audio         Pipewire, Pulseaudio and Bluetooth audio.
  cockpit       Cockpit web admin interface.
  crystalfontz  Crystalfontz LCD display.
  desktop       Desktop tools.
  dialout       The dialout group, which controls access to serial ports.
  filesharing   Configure Samba filesharing.
  login         Configure login and power settings.
  media         Basic media tools.
  mopidy        Mopidy media player.
  plusdeck      Plus Deck 2C PC Cassette Deck.
  ssh           SSH agent and user keys.
  stardeck      Stardeck software.
  starship      Starship shell prompt.
  vim           Vim text editor.
  shell         Hooks for the user shell.
```

## License

This software is licensed under the Mozilla Public License, Version 2.0.
