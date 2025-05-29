#!/usr/bin/env bash

set -euo pipefail

NGINX_TEMPLATE_DIR="$(pwd)/playbooks/web/templates"

set -x

prettier --write .

nginx-formatter \
  -char ' ' \
  -indent 4 \
  -input "${NGINX_TEMPLATE_DIR}" \
  -output "${NGINX_TEMPLATE_DIR}" \
