#!/usr/bin/env node

import process from 'node:process';

import minimist from 'minimist';

import {
  runAnsiblePlaybook,
  runParallelAnsiblePlaybooks,
  findAnsibleConfig,
  findStardeckConfig,
  logger,
  LOG_LEVELS,
  stages,
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

  async function ansible(stage, options) {
    const opts = {
      // TODO: Allow setting ansible verbosity separately
      logLevel,
      check,
      diff,
      varFiles: [configFile],
      configFile: ansibleConfigFile,
      serial,
      ...options,
    };
    if (typeof stage === 'string') {
      return runAnsiblePlaybook(stage, opts);
    }
    await runParallelAnsiblePlaybooks(stage, opts);
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

  // TODO: This is slow. Can we speed it up by either lifting packages or
  // using include_tasks instead of include_playbook?
  await ansible('packages.yml');

  //
  // Core playbooks
  //
  await ansible([
    { name: 'logind', playbook: 'core/logind.yml' },
    { name: 'sddm', playbook: 'core/sddm.yml' },
    { name: 'cockpit', playbook: 'cockpit/main.yml' },
    { name: 'ssh', playbook: 'ssh/main.yml' },
    { name: 'bluetooth', playbook: 'audio/bluetooth.yml' },
    { name: 'pipewire', playbook: 'audio/pipewire.yml' },
    { name: 'pulseaudio', playbook: 'audio/pulseaudio.yml' },
    { name: 'plusdeck', playbook: 'plusdeck.yml' },
    { name: 'dialout', playbook: 'dialout.yml' },
    { name: 'mopidy', playbook: 'mopidy/main.yml' },
    { name: 'filesharing', playbook: 'filesharing/main.yml' },
  ]);

  // TODO: Uses dnf and is flakey when run in parallel with other tasks.
  // This is unfortunate, because this one takes a really long time.
  await ansible('vim/main.yml');

  const devTasks = {
    git: { name: 'git', playbook: 'development/git.yml' },
    yadm: {
      name: 'yadm',
      playbook: 'development/yadm.yml',
      dependencies: ['git'],
    },
  };

  for (let stage of stages(devTasks)) {
    await ansible(stage);
  }
}

(async () => {
  try {
    await main();
  } catch (err) {
    logger.fatal(err);
  }
})();
