#!/usr/bin/env bash

set -euxo pipefail

# Modules to install from the local filesystem
WORKSPACE_MODULES=''

carton install

while IFS= read -r module; do
  carton run bash -c "cd ../${module} && dzil install"
done <<< "${WORKSPACE_MODULES}"
