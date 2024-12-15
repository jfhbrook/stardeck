#!/usr/bin/env bash

set -euo pipefail

DNF_PACKAGES=()
ASDF_PLUGINS=()

PACKAGE_PATH="${PACKAGE_PATH:-$(pwd)}"

if [[ -n "${1:-}" ]]; then
  while [[ $# -gt 0 ]]; do
    case "${1}" in
      --force)
        FORCE=1
        shift
        ;;
      *)
        source "${PACKAGE_PATH}/packages/${1}.sh"
        shift
        ;;
    esac
  done
else
  source "${PACKAGE_PATH}/package.sh"
  for package in ${ABSENT[@]}; do
    source "${PACKAGE_PATH}/packages/${package}.sh"
  done
fi

DNF_PACKAGES=($(printf "%s\n" "${DNF_PACKAGES[@]:-}" | sort -u))

while read -r function; do
  if [[ $function == "remove_"* ]]; then
    set +x
    $function
    set -x
  fi
done <<< $(typeset -F | sed 's/declare -f //')

if [ -n "${DNF_PACKAGES:-}" ]; then
  sudo dnf remove -y "${DNF_PACKAGES[@]}"
fi
