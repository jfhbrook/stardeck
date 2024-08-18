#!/usr/bin/env bash

sudo dnf install git

systemctl --user enable ssh-agent
systemctl --user start ssh-agent

git config --global user.name "Josh Holbrook"
git config --global user.email "josh.holbrook@gmail.com"
git config --global init.defaultBranch main
