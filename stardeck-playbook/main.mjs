#!/usr/bin/env node

import process from 'node:process';

import minimist from 'minimist';

import { ansiblePlaybookArgv, ansiblePlaybookEnv, findAnsibleConfig, INVENTORY_FILE, findStardeckConfig, loadStardeckConfig, VERBOSITY } from './index.mjs';

const HELP = `USAGE: stardeck-playbook OPTIONS

OPTIONS:
  -h|--help          Show this help text and exit.
  --feature FEATURE  Target a feature. May be specified more than once.
  --no-update        Do not run software updates.
`;

function main() {
  const argv = minimist(process.argv.slice(2), {
    alias: { h: 'help' },
    boolean: ['help', 'update'],
    string: ['feature'],
    default: {
      help: false,
      update: true,
    },
    alias: {
      h: 'help',
    },
    '--': true,
  });

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

  console.log(ansiblePlaybookArgv({
    inventoryFile: INVENTORY_FILE,
    verbosity: VERBOSITY.INFO,
    check: true,
    diff: true,
    askBecomePass: true,
    varFiles: [findStardeckConfig()],
  }));
}

main();
