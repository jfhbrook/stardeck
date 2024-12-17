#!/usr/bin/env bash

SCRIPT_PATH="$(realpath "${BASH_SOURCE[0]}")"
ANSIBLE_CONFIG="$(dirname "$(dirname "$(dirname "${SCRIPT_PATH}")")")/ansible.cfg"

export ANSIBLE_CONFIG

exec "$@"
