import * as path from 'node:path';
import { env } from 'node:process';

export const VERBOSITY = {
  DEBUG: 3,
  VERBOSE: 2,
  INFO: 1,
  WARNING: 0,
  ERROR: 0,
  FATAL: 0,
};

export function ansiblePlaybookArgv({
  inventoryFile,
  verbosity,
  check,
  diff,
  askBecomePass,
  varFiles,
  extraVars,
  tags,
  skipTags,
  listTags,
}) {
  const argv = ['-i', inventoryFile];

  if (typeof verbosity === 'number') {
    argv.push('-' + 'v'.repeat(verbosity));
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
