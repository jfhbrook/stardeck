#!/usr/bin/env bash

set -euxo pipefail

# TODO: Pass perlcritic at --cruel severity

find . \
  \( -name '*.pl' \
     -o -name '*.pm' \
  \) \
  ! -path './local/*' \
  ! -path '*/.build/*' \
  -print0 | xargs -0 perlcritic --harsh
