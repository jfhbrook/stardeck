import * as path from 'node:path';
import { env } from 'node:process';

import { INVENTORY_FILE, PLAYBOOK_DIR } from './config/index.mjs';
import { LOG_LEVELS } from './logging.mjs';

export const VERBOSITY = {
  DEBUG: 3,
  VERBOSE: 2,
  INFO: 1,
  WARNING: 0,
  ERROR: 0,
};

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
    argv.push('-' + 'v'.repeat(LOG_LEVELS[logLevel].verbosity));
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
  let envVars = {};
  Object.assign(envVars, {
    ANSIBLE_CONFIG: configFile,
  });
  Object.assign(envVars, env);
  return envVars;
}
