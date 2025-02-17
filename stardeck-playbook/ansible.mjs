import { spawnSync } from 'node:child_process';
import { createRequire } from 'node:module';
import * as path from 'node:path';
import { env } from 'node:process';

import {
  INVENTORY_FILE,
  PLAYBOOK_DIR,
  STARDECK_HOME,
} from './config/index.mjs';
import { logger, LOG_LEVELS } from './logging.mjs';

const require = createRequire(import.meta.url);

const quote = require('shell-quote/quote');

export function ansiblePlaybookArgv(
  playbook,
  {
    logLevel,
    check,
    diff,
    askBecomePass,
    varFiles,
    extraVars,
    tags,
    skipTags,
    listTags,
  },
) {
  const argv = ['-i', INVENTORY_FILE];

  if (logLevel) {
    const verbosity = LOG_LEVELS[logLevel].verbosity;
    if (verbosity) {
      argv.push('-' + 'v'.repeat(verbosity));
    }
  }

  if (check) {
    argv.push('--check');
  }

  if (diff) {
    argv.push('--diff');
  }

  if (askBecomePass) {
    argv.push('--ask-become-pass');
  }

  for (let varFile of varFiles || []) {
    argv.push('--extra-vars');
    argv.push(`@${path.resolve(varFile)}`);
  }

  if (extraVars) {
    argv.push('--extra-vars');
    argv.push(
      Object.entries(extraVars)
        .map(([name, value]) => `${name}=${value}`)
        .join(' '),
    );
  }

  if (tags) {
    argv.push('--tags');
    argv.push(tags.join(','));
  }

  if (skipTags) {
    argv.push('--skip-tags');
    argv.push(skipTags.join(','));
  }

  if (listTags) {
    argv.push('--list-tags');
  }

  argv.push(path.join(PLAYBOOK_DIR, playbook));

  return argv;
}

export function ansiblePlaybookEnv({ configFile }) {
  let envVars = {
    ANSIBLE_CONFIG: configFile,
  };

  let ansibleHome = `${process.env.HOME}/.ansible`;
  if (env.ANSIBLE_HOME && env.ANSIBLE_HOME.length) {
    ansibleHome = env.ANSIBLE_HOME;
  }
  envVars.ANSIBLE_HOME = ansibleHome;

  for (let [name, baseDir] of [
    ['ANSIBLE_ROLES_PATH', 'roles'],
    ['ANSIBLE_COLLECTIONS_PATH', 'collections'],
  ]) {
    envVars[name] = process.env[name] || '';
    if (!envVars[name].length) {
      envVars[name] = [
        `./${baseDir}`,
        path.join(STARDECK_HOME, baseDir),
        path.join(ansibleHome, baseDir),
        path.join('/usr/share/ansible', baseDir),
      ].join(':');
    } else if (!envVars[name].includes(dir)) {
      envVars[name] = [
        `./${baseDir}`,
        path.join(STARDECK_HOME, baseDir),
        envVars[name],
      ].join(':');
    }
  }

  return envVars;
}

export function runAnsiblePlaybook(playbook, options) {
  const command = ansiblePlaybookArgv(playbook, options);
  const env = { ...process.env, ...ansiblePlaybookEnv(options) };

  logger.debug(`Running ansible-playbook ${quote(command)}...`);

  const { status } = spawnSync('ansible-playbook', command, {
    env,
    stdio: 'inherit',
  });

  if (status) {
    logger.fatal(`ansible exited with status ${status}`);
  }
}

export function runSerialAnsiblePlaybooks(stage, globalOptions) {
  for (let { playbook, options } of stage) {
    runAnsiblePlaybook(playbook, {
      ...globalOptions,
      ...options,
    });
  }
}

export async function runParallelAnsiblePlaybooks(stage, globalOptions) {
  if (!stage.length) {
    return;
  }

  if (stage.length == 1) {
    return runAnsiblePlaybook(stage[0].playbook, {
      ...globalOptions,
      ...stage[0].options,
    });
  }

  if (globalOptions.serial) {
    return runSerialAnsiblePlaybooks(stage, globalOptions);
  }

  throw new Error('not implemented: runParallelAnsiblePlaybooks');
}
