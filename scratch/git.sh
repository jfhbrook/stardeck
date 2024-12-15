#!/usr/bin/env bash

set -euo pipefail

sudo dnf install git

git config --global user.name "Josh Holbrook"
git config --global user.email "josh.holbrook@gmail.com"
git config --global init.defaultBranch main
