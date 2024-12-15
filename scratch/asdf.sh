#!/usr/bin/env bash

set -euo pipefail

if [ ! -d ~/.asdf ]; then
  git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.14.1
fi

sudo dnf install openssl-devel libffi-devel libyaml-devel

. "$HOME/.asdf/asdf.sh"

asdf install ruby latest
asdf global ruby latest
