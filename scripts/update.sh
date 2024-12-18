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
# Run ansible. This installs the vast majority of things.
#

just -f ./playbooks/justfile playbook --ask-become-pass main.yml

git status
yadm status
