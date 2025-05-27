#!/usr/bin/env node

import path from 'node:path';
import process from 'node:process';

import minimist from 'minimist';

import {
  runAnsiblePlaybook,
  runParallelAnsiblePlaybooks,
  findAnsibleConfig,
  findStardeckConfig,
  loadStardeckConfig,
  logger,
  LOG_LEVELS,
} from './index.mjs';

const HELP = `USAGE: stardeck-playbook OPTIONS

OPTIONS:
  -h|--help              Show this help text and exit.
  --ansible-config FILE  The path to an ansible.cfg config file.
  --config-file FILE     The path to a stardeck.yml config file.
  --dry-run              Run ansible with --check and --diff.
  --feature FEATURE      Target a feature. May be specified more than once.
  --log-level LEVEL      Set the log level. Valid values are: ${Object.keys(LOG_LEVELS).join(', ')}
  --playbook PLAYBOOK    Run a specific playbook and exit.
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
  development   Development dependencies.
  dialout       The dialout group, which controls access to serial ports.
  filesharing   Configure Samba filesharing.
  login         Configure login and power settings.
  media         Basic media tools.
  mopidy        Mopidy media player.
  plusdeck      Plus Deck 2C PC Cassette Deck.
  ssh           SSH agent and user keys.
  shell         Hooks for the user shell.
  stardeck      Stardeck software.
  starship      Starship shell prompt.
  vim           Vim text editor.
  web           NGINX web server.

PLAYBOOKS:
  repositories.yml               System repositories.
  update.yml                     System updates.
  packages.yml                   System packages.
  audio/bluetooth.yml            Bluetooth audio configuration.
  audio/pipewire.yml             Pipewire configuration.
  audio/pulseaudio.yml           Pulseaudio configuration.
  cockpit/main.yml               Cockpit web admin interface.
  crystalfontz/main.yml          Crystalfontz LCD display.
  dialout.yml                    The dialout group, which controls access to
                                 serial ports.
  filesharing/main.yml           Configure Samba filesharing.
  login.yml                      Configure login and power settings.
  mopidy/main.yml                Mopidy media player.
  plusdeck/main.yml              Plus Deck 2C PC Cassette Deck.
  ssh/main.yml                   SSH agent and user keys.
  starship/main.yml              Starship shell prompt.
  vim/main.yml                   Vim text editor.
  web/main.yml                   NGINX configuration.
  development/git.yml            Git configuration.
  development/gomplate.yml       Gomplate CLI template generator.
  development/hugo.yml           Hugo static site generator.
  development/neovim.yml         Neovim text editor.
  development/node-dev/main.yml  Node.js development environment.
  development/perl-dev/main.yml  Perl development environment.
  development/rust-dev/main.yml  Rust development environment.
 
`;

export async function main() {
  const argv = minimist(process.argv.slice(2), {
    boolean: ['help', 'dry-run', 'serial', 'update'],
    string: ['ansible-config', 'config-file', 'feature', 'log-level'],
    default: {
      help: false,
      'ansible-config': null,
      'config-file': null,
      'dry-run': false,
      'log-level': 'warn',
      serial: false,
      update: true,
    },
    alias: {
      h: 'help',
    },
  });

  const logLevel = argv['log-level'];

  try {
    logger.setLevel(argv['log-level']);
  } catch (exc) {
    logger.fatal(exc.message);
  }

  if (argv.help) {
    console.log(HELP);
    process.exit(0);
  }

  const playbook = argv.playbook;

  let features = argv.feature;
  if (typeof features === 'undefined') {
    features = [];
  } else if (typeof features === 'string') {
    features = [features];
  }

  const check = argv['dry-run'];
  const diff = check;
  const update = argv.update;
  const serial = argv.serial;

  let tags = Array.from(features);

  if (features.length && update) {
    tags.push('update');
  }

  let skipTags = undefined;

  if (!features.length && !update) {
    skipTags = ['update'];
  }

  let configFile = argv['config-file'];
  if (!configFile || !configFile.length) {
    configFile = findStardeckConfig();
  }

  const config = loadStardeckConfig(configFile);

  let ansibleConfigFile = argv['ansible-config'];
  if (!ansibleConfigFile || !ansibleConfigFile.length) {
    ansibleConfigFile = findAnsibleConfig();
  }

  const stardeckConfigHome = path.dirname(configFile);
  const ansibleConfigHome = path.dirname(ansibleConfigFile);

  function enabled(feature) {
    return !features.length || features.includes(feature);
  }

  async function ansible(stage, options) {
    const opts = {
      // TODO: Allow setting ansible verbosity separately
      logLevel,
      check,
      diff,
      varFiles: [configFile],
      extraVars: {
        stardeck_config_home: stardeckConfigHome,
        ansible_config_home: ansibleConfigHome,
        ...(options?.extraVars || {}),
      },
      configFile: ansibleConfigFile,
      serial,
      tags,
      skipTags,
      ...options,
    };
    if (typeof stage === 'string') {
      return runAnsiblePlaybook(stage, opts);
    }
    await runParallelAnsiblePlaybooks(
      stage.filter((stg) => !stg.feature || enabled(stg.feature)),
      opts,
    );
  }

  //
  // Given a specific playbook, run it and exit.
  //
  if (playbook) {
    await ansible(playbook);
    return;
  }

  //
  // Set up expected repositories
  //

  await ansible('repositories.yml');

  //
  // Run global updates
  //

  if (update) {
    await ansible('update.yml');
  }

  //
  // Install global packages
  //

  await ansible('packages.yml');

  //
  // Core playbooks
  //
  await ansible([
    { feature: 'audio', name: 'bluetooth', playbook: 'audio/bluetooth.yml' },
    { feature: 'audio', name: 'pipewire', playbook: 'audio/pipewire.yml' },
    { feature: 'audio', name: 'pulseaudio', playbook: 'audio/pulseaudio.yml' },
    { feature: 'cockpit', name: 'cockpit', playbook: 'cockpit/main.yml' },
    {
      feature: 'crystalfontz',
      name: 'crystalfontz',
      playbook: 'crystalfontz/main.yml',
    },
    { feature: 'dialout', name: 'dialout', playbook: 'dialout.yml' },
    {
      feature: 'filesharing',
      name: 'filesharing',
      playbook: 'filesharing/main.yml',
    },
    { feature: 'login', name: 'login', playbook: 'login.yml' },
    { feature: 'mopidy', name: 'mopidy', playbook: 'mopidy/main.yml' },
    { feature: 'plusdeck', name: 'plusdeck', playbook: 'plusdeck/main.yml' },
    { feature: 'ssh', name: 'ssh', playbook: 'ssh/main.yml' },
    { feature: 'starship', name: 'starship', playbook: 'starship/main.yml' },
    { feature: 'vim', name: 'vim', playbook: 'vim/main.yml' },
    { feature: 'web', name: 'web', playbook: 'web/main.yml' },
  ]);

  if (config.development && enabled('development')) {
    await ansible([
      { name: 'git', playbook: 'development/git.yml' },
      { name: 'gomplate', playbook: 'development/gomplate.yml' },
      { name: 'hugo', playbook: 'development/hugo.yml' },
      { name: 'neovim', playbook: 'development/neovim.yml' },
      { name: 'node', playbook: 'development/node-dev/main.yml' },
      { name: 'perl', playbook: 'development/perl-dev/main.yml' },
      { name: 'rust', playbook: 'development/rust-dev/main.yml' },
    ]);
  }

  //
  // ~/.bashrc
  //

  if (enabled('shell')) {
    await ansible('shell.yml');
  }
}

(async () => {
  try {
    await main();
  } catch (err) {
    logger.fatal(err);
  }

  logger.info('ok');
})();
