#!/usr/bin/env bash

set -euo pipefail

sudo dnf install https://mirrors.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm https://mirrors.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm

# NOTE: Different command on Fedora 41
sudo dnf config-manager --enable fedora-cisco-openh264
