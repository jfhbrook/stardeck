#!/usr/bin/env bash

find . -maxdepth 1 -name '*.yml' ! -name 'requirements.yml' ! -name 'inventory.yml' -print0 | xargs -0 "$@"
