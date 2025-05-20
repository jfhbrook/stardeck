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
  set -x
  sudo dnf install -y "${DNF_PACKAGES[@]}"
fi
