#!/usr/bin/env bash

set -euo pipefail

kpackagetool6 --type=KWin/Script -i .

kwriteconfig6 --file kwinrc --group Plugins --key myscriptEnabled true

dbus-send --session --dest=org.kde.KWin --print-reply /KWin reconfigure
