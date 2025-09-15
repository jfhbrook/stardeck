#!/usr/bin/env bash

set -euxo pipefail

dbus-send --session --dest=org.kde.KWin --print-reply /KWin org.kde.KWin.reconfigure
