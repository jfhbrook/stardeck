#!/usr/bin/env bash

import process from 'node:process';

import minimist from 'minimist';

import { findAnsibleConfig, runAnsibleGalaxyInstall } from '../index.mjs';

const HELP = `USAGE: ./scripts/install.mjs OPTIONS

OPTIONS:
  -h|--help              Show this help text and exit.
  --ansible-config FILE  The path to an ansible.cfg config file

ENVIRONMENT:
  ANSIBLE_CONFIG         A path to an ansible.cfg configuration file.
  STARDECK_CONFIG_HOME   A directory containing stardeck configuration files.
`;

export async function main() {
  const argv = minimist(process.argv.slice(2), {
    boolean: ['help'],
    string: ['ansible-config'],
    default: {
      help: false,
      'ansible-config': null,
    },
    alias: {
      h: 'help',
    },
  });

  if (argv.help) {
    console.log(HELP);
    process.exit(0);
  }

  let configFile = argv['ansible-config'];
  if (!configFile || !configFile.length) {
    configFile = findAnsibleConfig();
  }

  const requirementsFile = argv._.length
    ? argv._.shift()
    : './requirements.yml';
  const ansibleArgv = argv._;

  runAnsibleGalaxyInstall(requirementsFile, { configFile, ansibleArgv });
}

(async () => {
  try {
    await main();
  } catch (err) {
    console.error(err);
    process.exit(1);
  }
})();
