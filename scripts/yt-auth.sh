#!/usr/bin/env bash

set -euo pipefail

if [ ! -d /etc/mopidy/ytmusic ]; then
  sudo mkdir -p /etc/mopidy/ytmusic
fi

if [ ! -f /etc/mopidy/ytmusic/auth.json ]; then
  (cd /etc/mopidy/ytmusic && sudo mopidyctl ytmusic setup)
else
  (cd /etc/mopidy/ytmusic && sudo mopidyctl ytmusic reauth)
fi
