#!/usr/bin/env bash

set -euxo pipefail

find . \
  \( -name '*.pl' \
     -o -name '*.pm' \
     -o -name 'cpanfile' \
  \) \
  ! -path './local/*' \
  -print0 | xargs -0 perltidy -b -bext='/'
