#!/usr/bin/env node

import process from 'node:process';

import minimist from 'minimist';

import {
  ansiblePlaybookArgv,
  ansiblePlaybookEnv,
  findAnsibleConfig,
  INVENTORY_FILE,
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

ENVIRONMENT:
  ANSIBLE_CONFIG         A path to an ansible.cfg configuration file.
  STARDECK_CONFIG_FILE   A path to a stardeck.yml configuration file.
  STARDECK_CONFIG_HOME   A directory containing stardeck configuration files.
`;

function main() {
  // TODO: What env vars does ansible support?
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

  // let update = argv.update;

  let configFile = argv['config-file'];
  if (!configFile || !configFile.length) {
    configFile = findStardeckConfig();
  }

  let ansibleConfigFile = argv['ansible-config'];
  if (!ansibleConfigFile || !ansibleConfigFile.length) {
    ansibleConfigFile = findAnsibleConfig();
  }

  function ansibleArgv(playbook, options) {
    return ansiblePlaybookArgv(playbook, {
      ...options,
      logLevel,
      check,
      diff,
      askBecomePass: false,
      varFiles: [configFile],
    });
  }

  const ansibleEnv = ansiblePlaybookEnv({ configFile: ansibleConfigFile });
  console.log(ansibleEnv);
  console.log(ansibleArgv('repositories.yml'));
  console.log(ansibleArgv('update.yml'));
}

main();
