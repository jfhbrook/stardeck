#!/usr/bin/env bash

set -euxo pipefail

cp ./systemd/* ~/.config/systemd/user/

systemctl --user enable stardeck.service
systemctl --user start stardeck.service
