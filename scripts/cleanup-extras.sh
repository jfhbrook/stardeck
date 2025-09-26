#!/usr/bin/env bash

#
# Adopted from a Fedora kernel cleanup script from:
#
#     https://docs.fedoraproject.org/en-US/quick-docs/upgrading-fedora-offline/
#

sudo dnf list --extras | grep -v watchexec-cli

read -p 'Would you like to remove extras? [y/N]' -n 1 -r
echo
if [[ "${REPLY}" =~ ^[Yy]$ ]]; then
  # shellcheck disable=SC2046
  sudo dnf remove $(sudo dnf repoquery --extras --exclude=kernel,kernel-\*,kmod-\*,watchexec-cli)
  sudo dnf autoremove
else
  # shellcheck disable=SC2016
  echo '
To remove all extras, run:
    sudo dnf remove $(sudo dnf repoquery --extras --exclude=kernel,kernel-\*,kmod-\*)
'
fi
