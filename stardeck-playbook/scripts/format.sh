#!/usr/bin/env bash

set -euxo pipefail

prettier --write .

nginx-formatter \
  -char ' ' \
  -indent 4 \
  -input "$(pwd)/playbooks/web/templates" \
  -output "$(pwd)/playbooks/web/templates"
