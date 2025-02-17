#!/usr/bin/env bash

set -euo pipefail

#
# Basic boostrap. Without these, updates will fail.
#

DNF_PACKAGES=()

if ! which ansible > /dev/null; then
  DNF_PACKAGES+=(ansible)
fi

if ! which just > /dev/null; then
  DNF_PACKAGES+=(just)
fi

if [ -n "${DNF_PACKAGES:-}" ]; then
  sudo dnf install -y "${DNF_PACKAGES[@]}"
fi

#
# Run stardeck-playbook.
#

STARDECK_CONFIG_HOME="$(pwd)"
export STARDECK_CONFIG_HOME

(cd ./stardeck-playbook && sudo node main.mjs)

#
# Run legacy ansible playbooks.
#

just -f ./playbooks/justfile playbook --ask-become-pass "$@" main.yml

git status
yadm status
