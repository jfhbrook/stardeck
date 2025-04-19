#!/usr/bin/env bash

set -euxo pipefail

find . \
  \( -name '*.pl' \
     -o -name '*.pm' \
     -o -name '*.t' \
     -o -name 'cpanfile' \
  \) \
  ! -path './local/*' \
  ! -path '*/.build/*' \
  -print0 | xargs -0 perltidy -b -bext='/'
