#!/usr/bin/env bash

set -euxo pipefail

find . -name 'dist.ini' \
  ! -path './local/*' \
  -execdir dzil test \;
