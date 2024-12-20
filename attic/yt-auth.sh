#!/usr/bin/env bash

#
# This hacky script was built for authenticating mopidy-ytmusic. Today, this
# mopidy plugin is super busted, so I don't need to keep it around.
#
# If I do decide to dust off mopidy-ytmusic, it could very well be worth
# trying to pull authentication from yt-dl or yt-dlp - a suggestion from Tom K.
#

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
