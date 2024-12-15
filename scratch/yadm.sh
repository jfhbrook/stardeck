#!/usr/bin/env bash

set -euo pipefail

sudo dnf config-manager --add-repo https://download.opensuse.org/repositories/home:TheLocehiliosan:yadm/Fedora_40/home:TheLocehiliosan:yadm.repo
sudo dnf install yadm
