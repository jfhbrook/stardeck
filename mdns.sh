#!/usr/bin/env bash

sudo dnf install nss-mdns avahi

sudo systemctl enable --now avahi-daemon.service
