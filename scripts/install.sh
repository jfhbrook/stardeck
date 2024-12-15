#!/usr/bin/env bash

set -euo pipefail

DNF_PACKAGES=()
ASDF_PLUGINS=()

FORCE=""

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
  for package in ${PRESENT[@]}; do
    source "${PACKAGE_PATH}/packages/${package}.sh"
  done
fi

DNF_PACKAGES=($(printf "%s\n" "${DNF_PACKAGES[@]-}" | sort -u))

if [ -n "${DNF_PACKAGES:-}" ]; then
  sudo dnf install -y "${DNF_PACKAGES[@]}"
fi

if [ -n "${ASDF_PLUGINS:-}" ]; then
  for plugin in "${ASDF_PLUGINS[@]}"; do
    if ! asdf plugin list | grep -q "${plugin}"; then
      echo "ASDF plugin install not automated."
      echo "ASDF plugin ${plugin} must be installed manually."
    fi
  done
fi

while read -r function; do
  if [[ $function == "install_"* ]]; then
    set -x
    $function
    set +x
  fi
done <<< $(typeset -F | sed 's/declare -f //')
