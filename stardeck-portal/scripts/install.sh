#!/usr/bin/env bash

set -euxo pipefail

mkdir -p /srv/portal
rsync -a --progress --delete ./public/ /srv/portal
