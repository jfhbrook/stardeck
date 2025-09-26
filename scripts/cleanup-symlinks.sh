#!/usr/bin/env bash

sudo symlinks -r /usr/ | grep dangling

echo
read -p 'Would you like to remove dangling symlinks? [y/N]' -n 1 -r
echo
if [[ "${REPLY}" =~ ^[Yy]$ ]]; then
  sudo symlinks -r -d /usr
else
  echo '
To remove all dangling symlinks, run:
    sudo symlinks -r -d /usr
'
fi
