#!/usr/bin/env bash

set -euxo pipefail

STATIC_ROOT=/srv/portal
STATIC_OWNER=root
STATIC_GROUP=root

mkdir -p "${STATIC_ROOT}"
chown "${STATIC_OWNER}:${STATIC_GROUP}" "${STATIC_ROOT}"

rsync -a \
  --chown="${STATIC_OWNER}:${STATIC_GROUP}" \
  --progress \
  --delete \
  ./public/ "${STATIC_ROOT}"

semanage fcontext \
  --add \
  --type httpd_sys_content_t \
  "${STATIC_ROOT}(/.*)?"

restorecon -R "${STATIC_ROOT}"
