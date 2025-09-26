#!/usr/bin/env bash

for repo in $(dnf copr list | grep -v ' (disabled)$'); do
  sudo dnf copr remove "${repo}"
  sudo dnf clean all
  sudo dnf copr enable "${repo}"
done
