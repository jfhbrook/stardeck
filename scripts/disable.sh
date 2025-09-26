#!/usr/bin/env bash

set -euxo pipefail

systemctl --user stop stardeck.service
systemctl --user disable stardeck.service

for file in ./systemd/user/*; do
  rm ~/.config/systemd/user/"$(basename "${file}")"
done
