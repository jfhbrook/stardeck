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

just -f ./playbooks/justfile playbook main.yml

#
# Install rust. The rust installer is interactive, so we can't easily do it in
# ansible.
#

if ! which rust-analyzer > /dev/null; then
  sudo dnf install -y rust-analyzer
fi

if [ ! -d ~/.cargo ]; then
  curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
fi

for toolchain in stable nightly; do
  (. "${HOME}/.cargo/env" && rustup update "${toolchain}") || exit 1
done

git status
yadm status
