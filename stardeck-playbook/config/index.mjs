import * as fs from 'node:fs';
import { homedir } from 'node:os';
import * as path from 'node:path';
import { env } from 'node:process';

import yaml from 'yaml';

import { logger } from '../logging.mjs';

let XDG_CONFIG_DIR = path.join(homedir(), '.config');

if (env.XDG_CONFIG_DIR) {
  XDG_CONFIG_DIR = env.XDG_CONFIG_DIR;
} else if (env.XDG_CONFIG_DIRS) {
  XDG_CONFIG_DIR = env.XDG_CONFIG_DIRS.split(':')[0];
}

export { XDG_CONFIG_DIR };

export const STARDECK_HOME = path.dirname(
  path.dirname(import.meta.url.replace(/^file:/, '')),
);

export const GLOBAL_CONFIG_DIR = '/etc/stardeck';
export const LOCAL_CONFIG_DIR = path.join(XDG_CONFIG_DIR, 'stardeck');
export const DEFAULT_CONFIG_DIR = path.join(STARDECK_HOME, 'config');
export const PLAYBOOK_DIR = path.join(STARDECK_HOME, 'playbooks');

export function localConfigPath(filename) {
  return path.join(LOCAL_CONFIG_DIR, filename);
}

export function globalConfigPath(filename) {
  return path.join(GLOBAL_CONFIG_DIR, filename);
}

export function defaultConfigPath(filename) {
  return path.join(DEFAULT_CONFIG_DIR, filename);
}

export function readYamlFile(filename) {
  const contents = fs.readFileSync(filename, 'utf8');
  return yaml.parse(contents);
}

const FOUND_FILES = {};

export function findConfigFile(filename) {
  if (!FOUND_FILES[filename]) {
    logger.debug(`Searching for ${filename}...`);
    try {
      fs.accessSync(localConfigPath(filename), fs.constants.R_OK);
      FOUND_FILES[filename] = localConfigPath(filename);
    } catch (err) {
      logger.debug(err.message);
      try {
        fs.accessSync(globalConfigPath(filename), fs.constants.R_OK);
        FOUND_FILES[filename] = globalConfigPath(filename);
      } catch (err) {
        logger.debug(err.message);
        fs.accessSync(defaultConfigPath(filename));
        FOUND_FILES[filename] = defaultConfigPath(filename);
      }
    }
    logger.debug(`Found ${filename} at ${FOUND_FILES[filename]}`);
  }
  return FOUND_FILES[filename];
}

export function loadConfigFile(reader, filename) {
  return reader(findConfigFile(filename));
}

export function loadStardeckConfig(filename) {
  if (filename) {
    return readYamlFile(filename);
  }
  return loadConfigFile(readYamlFile, 'stardeck.yml');
}

export function findStardeckConfig(filename) {
  if (filename) {
    return filename;
  }
  return findConfigFile('stardeck.yml');
}

export function findAnsibleConfig(filename) {
  if (filename) {
    return filename;
  }
  return findConfigFile('ansible.cfg');
}

export const INVENTORY_FILE = path.join(DEFAULT_CONFIG_DIR, 'inventory.yml');
