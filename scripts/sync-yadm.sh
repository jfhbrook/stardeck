#!/usr/bin/env bash

set -euxo pipefail

REPO="$(yq .users.josh.yadm_repo < stardeck.yml)"

if [ ! -d ~/.local/share/yadm/repo.git ]; then
  yadm clone "${REPO}"
else
  yadm pull;
fi
