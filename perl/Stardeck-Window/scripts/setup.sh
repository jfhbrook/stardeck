#!/usr/bin/env bash

set -euxo pipefail

# Modules to install from the local filesystem
WORKSPACE_MODULES='Stardeck-Process'

carton install

while IFS= read -r module; do
  carton run bash -c "cd ../${module} && dzil install"
done <<< "${WORKSPACE_MODULES}"

carton run dzil install
