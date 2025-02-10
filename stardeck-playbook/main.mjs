#!/usr/bin/env node

import process from 'node:process';

import minimist from 'minimist';

const HELP = `USAGE: stardeck-playbook OPTIONS

COMMANDS:
  config  Configure
  update  Update
`;

function main() {
  const argv = minimist(process.argv.slice(2), {
    alias: { h: 'help' },
    boolean: ['help'],
    default: {
      help: false,
    },
  });

  if (!argv._.length) {
    console.log(HELP);
    process.exit(1);
  }

  if (argv.help) {
    console.log(HELP);
    process.exit(0);
  }
}

main();
