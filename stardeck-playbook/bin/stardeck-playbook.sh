#!/usr/bin/env bash

# TODO: Write this in perl - this is just a sketch of the general shape.

set -euo pipefail

# Constants. Can be overridden in build if need be.
STARDECK_GLOBAL_CONFIG_HOME=/etc/stardeck
STARDECK_LOCAL_CONFIG_HOME=~/.config/stardeck

# TODO: Set this based on CLI arguments, default if root user
STARDECK_CONFIG_HOME="${STARDECK_CONFIG_HOME:-${STARDECK_GLOBAL_CONFIG_HOME}}"
STARDECK_BIN_PATH="$(realpath "${BASH_SOURCE[0]}")"
STARDECK_SHARE_HOME="${STARDECK_SHARE_HOME:-$(dirname "${STARDECK_BIN_PATH}")}"

# Config files
STARDECK_CONFIG="${STARDECK_CONFIG_HOME}/stardeck.yml"
ANSIBLE_CONFIG="${STARDECK_CONFIG_HOME}/ansible.cfg"

# Share files
INVENTORY_FILE="${STARDECK_SHARE_HOME}/inventory.yml"

# TODO: Load config

# TODO: Run different playbooks with different settings
# https://stackoverflow.com/questions/54389215/run-multiple-jobs-within-perl-script-at-the-same-time
ANSIBLE_CONFIG="${ANSIBLE_CONFIG}" ansible-playbook -i "${INVENTORY_FILE}" "$@"
