#!/usr/bin/env bash

set -euo pipefail

# DNF4 installation commands
# (different in f41)
sudo dnf install 'dnf-command(config-manager)'
sudo dnf config-manager --add-repo https://cli.github.com/packages/rpm/gh-cli.repo
sudo dnf install gh --repo gh-cli
