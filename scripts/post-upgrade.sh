#!/usr/bin/env bash

#
# See: https://docs.fedoraproject.org/en-US/quick-docs/upgrading-fedora-offline/
#

# Update system configuration files
sudo rpmconf -a

# Upgrade COPR repos
# TODO: This is probably not necessary. But we can check this next upgrade.
# bash ./scripts/reconfigure-copr-repos.sh

# Run package updates
just update

# Remove retired packages
sudo remove-retired-packages

# Remove duplicate packages
sudo dnf remove --duplicates

# Remove unused packages
sudo dnf autoremove

# Clean up extra packages
just cleanup-extras

# Clean up old kernels
just cleanup-kernels

# Clean up old GPG pubkeys
sudo clean-rpm-gpg-pubkey

# Clean up dangling symlinks
just cleanup-symlinks
