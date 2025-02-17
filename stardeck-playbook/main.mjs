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

ENVIRONMENT:
  ANSIBLE_CONFIG         A path to an ansible.cfg configuration file.
  STARDECK_CONFIG_FILE   A path to a stardeck.yml configuration file.
  STARDECK_CONFIG_HOME   A directory containing stardeck configuration files.
`;

export async function main() {
  const argv = minimist(process.argv.slice(2), {
    boolean: ['help', 'dry-run', 'update'],
    string: ['ansible-config', 'config-file', 'feature', 'log-level'],
    default: {
      help: false,
      'ansible-config': null,
      'config-file': null,
      'dry-run': false,
      'log-level': 'warn',
      update: true,
    },
    alias: {
      h: 'help',
    },
    '--': true,
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
      serial: true,
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

  await ansible('repositories/main.yml');

  //
  // Run global updates
  //

  if (update) {
    await ansible('update.yml');
  }

  //
  // Core playbooks
  //
  await ansible([
    { name: 'logind', playbook: 'core/logind.yml' },
    { name: 'sddm', playbook: 'core/sddm.yml' },
    { name: 'wifi', playbook: 'core/wifi.yml' },
    { name: 'yt-dlp', playbook: 'core/yt-dlp.yml' },
    { name: 'cockpit', playbook: 'cockpit/main.yml' },
    { name: 'ssh', playbook: 'ssh/main.yml' },
    { name: 'vim', playbook: 'vim/main.yml' },
  ]);
}

main();
