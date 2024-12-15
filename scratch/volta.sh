#!/usr/bin/env bash

set -euo pipefail

if [ ! -d ~/.volta ]; then
  curl https://get.volta.sh | bash
fi

export VOLTA_HOME="$HOME/.volta"
export PATH="$VOLTA_HOME/bin:$PATH"

volta install node
