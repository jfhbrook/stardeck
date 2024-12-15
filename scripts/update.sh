#!/usr/bin/env bash

set -euo pipefail

PRESENT=()
ABSENT=()
DNF_PACKAGES=()
ASDF_PLUGINS=()

PACKAGE_PATH="${PACKAGE_PATH:-$(pwd)}"

source "${PACKAGE_PATH}/package.sh"

for package in ${PRESENT[@]}; do
  source "${PACKAGE_PATH}/packages/${package}.sh"
done

set -x
sudo dnf -y update
asdf update
set +x

ASDF_VERSIONS=''
if [ -n "${ASDF_PLUGINS:-}" ]; then
  for plugin in "${ASDF_PLUGINS}"; do
    OLD_VERSION=$(cd ~ && asdf current "${plugin}" | awk '{print $2}')

    set -x
    asdf install "${plugin}" latest
    asdf global "${plugin}" latest
    set +x

    NEW_VERSION=$(cd ~ && asdf current "${plugin}" | awk '{print $2}')

    if [[ "${OLD_VERSION}" != "${NEW_VERSION}" ]]; then
      if [ -n "${ASDF_VERSIONS}" ]; then
        ASDF_VERSIONS="${ASDF_VERSIONS}, "
      fi
      ASDF_VERSIONS="${plugin} ${NEW_VERSION}"
    fi
  done
fi

while read -r function; do
  if [[ $function == "update_"* ]]; then
    set -x
    $function
    set +x
  fi
done <<< $(typeset -F | sed 's/declare -f //')

if [ -n "${ASDF_VERSIONS}" ]; then
  set -x
  yadm add ~/.tool-versions
  yadm commit -m "${ASDF_VERSIONS}"
  yadm push || echo "WARN: Failed to push asdf version changes"
  set +x
fi

set -x
yadm status
