#!/usr/bin/env bash

set -euxo pipefail

kwriteconfig6 \
  --file kwinrc \
  --group Plugins \
  --key stardeck-kwinEnabled \
  false

bash ./scripts/reconfigure.sh
