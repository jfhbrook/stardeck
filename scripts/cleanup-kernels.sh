#!/usr/bin/env bash

#
# Adopted from a Fedora kernel cleanup script from:
#
#     https://docs.fedoraproject.org/en-US/quick-docs/upgrading-fedora-offline/
#

# shellcheck disable=SC2207
old_kernels=($(dnf repoquery --installonly --latest-limit=-1 -q))

if [ "${#old_kernels[@]}" -eq 0 ]; then
    echo "No old kernels found"
    exit
fi

dnf remove "${old_kernels[@]}"
