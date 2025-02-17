#!/usr/bin/env bash

set -euo pipefail

STARDECK_CONFIG_HOME="$(pwd)"
export STARDECK_CONFIG_HOME

cd ./stardeck-playbook && sudo -E node main.mjs "$@"
