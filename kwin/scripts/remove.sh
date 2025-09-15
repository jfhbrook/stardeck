#!/usr/bin/env bash

set -euxo pipefail

bash ./scripts/disable.sh

kpackagetool6 --type=KWin/Script -r stardeck-kwin
