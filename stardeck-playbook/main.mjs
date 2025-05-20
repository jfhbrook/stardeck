#!/usr/bin/env node

import path from 'node:path';
import process from 'node:process';

import minimist from 'minimist';

import {
  runAnsiblePlaybook,
  runParallelAnsiblePlaybooks,
  findAnsibleConfig,
  findStardeckConfig,
  logger,
  LOG_LEVELS,
} from './index.mjs';

const HELP = `USAGE: stardeck-playbook OPTIONS

OPTIONS:
  -h|--help              Show this help text and exit.
  --ansible-config FILE  The path to an ansible.cfg config file
  --config-file FILE     The path to a stardeck.yml config file
  --dry-run              Run ansible with --check and --diff
  --feature FEATURE      Target a feature. May be specified more than once.
  --log-level LEVEL      Set the log level. Valid values are: ${Object.keys(LOG_LEVELS).join(', ')}
  --no-update            Do not run software updates.
  --serial               Run playbooks in order

ENVIRONMENT:
  ANSIBLE_CONFIG         A path to an ansible.cfg configuration file.
  STARDECK_CONFIG_FILE   A path to a stardeck.yml configuration file.
  STARDECK_CONFIG_HOME   A directory containing stardeck configuration files.

FEATURES:
  core
  audio
  cockpit
  crystalfontz
  desktop
  dialout
  filesharing
  mopidy
  plusdeck
  ssh
  stardeck
  vim
  shell
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

  let features = argv.feature;
  if (typeof features === 'undefined') {
    features = [];
  } else if (typeof features === 'string') {
    features = [features];
  }

  const tags = features.length ? features.concat('update') : [];

  const check = argv['dry-run'];
  const diff = check;
  const update = argv.update;
  const serial = argv.serial;

  let configFile = argv['config-file'];
  if (!configFile || !configFile.length) {
    configFile = findStardeckConfig();
  }

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
    { name: 'vim', playbook: 'vim/main.yml', feature: 'vim' },
    { name: 'logind', playbook: 'core/logind.yml', feature: 'core' },
    { name: 'sddm', playbook: 'core/sddm.yml', feature: 'core' },
    { name: 'cockpit', playbook: 'cockpit/main.yml', feature: 'cockpit' },
    { name: 'ssh', playbook: 'ssh/main.yml', feature: 'ssh' },
    { name: 'bluetooth', playbook: 'audio/bluetooth.yml', feature: 'audio' },
    { name: 'pipewire', playbook: 'audio/pipewire.yml', feature: 'audio' },
    { name: 'pulseaudio', playbook: 'audio/pulseaudio.yml', feature: 'audio' },
    { name: 'plusdeck', playbook: 'plusdeck/main.yml', feature: 'plusdeck' },
    {
      name: 'crystalfontz',
      playbook: 'crystalfontz/main.yml',
      feature: 'crystalfontz',
    },
    { name: 'dialout', playbook: 'dialout.yml', feature: 'dialout' },
    { name: 'mopidy', playbook: 'mopidy/main.yml', feature: 'mopidy' },
    {
      name: 'filesharing',
      playbook: 'filesharing/main.yml',
      feature: 'filesharing',
    },
  ]);

  if (configFile.development && !features.length) {
    await ansible([
      { name: 'git', playbook: 'development/git.yml' },
      { name: 'gomplate', playbook: 'development/gomplate.yml' },
      { name: 'neovim', playbook: 'development/neovim.yml' },
      { name: 'node', playbook: 'development/node-dev/main.yml' },
      { name: 'perl', playbook: 'development/perl-dev/main.yml' },
      { name: 'rust', playbook: 'development/rust-dev/main.yml' },
      { name: 'starship', playbook: 'development/starship/main.yml' },
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
