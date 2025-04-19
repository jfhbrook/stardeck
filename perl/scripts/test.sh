#!/usr/bin/env bash

set -euxo pipefail

find . -name 'dist.ini' \
  ! -path './local/*' \
  ! -path '*/.build/*' \
  -execdir carton run dzil test \;
