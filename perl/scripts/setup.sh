#!/usr/bin/env bash

set -euxo pipefail

carton install

find . -name 'dist.ini' \
  ! -path './local/*' \
  ! -path '*/.build/*' \
  -execdir bash ./scripts/setup.sh \;
