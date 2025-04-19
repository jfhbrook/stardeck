#!/usr/bin/env bash

set -euxo pipefail

if [ ! -d local ]; then
  carton install
fi

# find . -name 'dist.ini' \
#   ! -path './local/*' \
#   ! -path '*/.build/*' \
#   -execdir pwd \;
