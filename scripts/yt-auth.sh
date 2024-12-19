#!/usr/bin/env bash

set -euo pipefail

if [ ! -d /etc/mopidy/ytmusic ]; then
  sudo mkdir -p /etc/mopidy/ytmusic
fi

sudo rm -f /etc/mopidy/ytmusic/auth.json
sudo rm -f /tmp/auth.json
(cd /tmp && sudo mopidyctl ytmusic setup)
sudo cp /tmp/auth.json /etc/mopidy/ytmusic/auth.json
sudo chown mopidy:mopidy /etc/mopidy/ytmusic/auth.json
sudo rm /tmp/auth.json
