#!/usr/bin/env bash

set -euxo pipefail

kpackagetool6 --type=KWin/Script -i .

bash ./scripts/enable.sh
