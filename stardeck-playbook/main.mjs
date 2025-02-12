#!/usr/bin/env node

import process from 'node:process';

import minimist from 'minimist';

import {
  ansiblePlaybookArgv,
  ansiblePlaybookEnv,
  findAnsibleConfig,
  INVENTORY_FILE,
  findStardeckConfig,
  loadStardeckConfig,
  logger,
  LOG_LEVELS,
} from './index.mjs';

const HELP = `USAGE: stardeck-playbook OPTIONS

OPTIONS:
  -h|--help          Show this help text and exit.
  --feature FEATURE  Target a feature. May be specified more than once.
  --no-update        Do not run software updates.
`;

function main() {
  const argv = minimist(process.argv.slice(2), {
    boolean: ['help', 'update'],
    string: ['feature', 'log-level'],
    default: {
      help: false,
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

  let features = argv.feature;
  if (typeof features === 'undefined') {
    features = [];
  } else if (typeof features === 'string') {
    features = [features];
  }

  let update = argv.update;
  const config = loadStardeckConfig();

  if (argv.help) {
    console.log(HELP);
    process.exit(0);
  }

  logger.warning(
    ansiblePlaybookArgv('main.yml', {
      inventoryFile: INVENTORY_FILE,
      logLevel,
      check: true,
      diff: true,
      askBecomePass: true,
      varFiles: [findStardeckConfig()],
    }),
  );
}

main();
