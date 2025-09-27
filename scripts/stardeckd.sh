#!/usr/bin/env bash

STARDECK_HOME="$(realpath "$(dirname "$(dirname "${BASH_SOURCE[0]}")")")"

cd "${STARDECK_HOME}" || exit 1

exec go run ./cmd/stardeckd/main.go "$@"
