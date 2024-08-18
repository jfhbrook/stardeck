#!/usr/bin/env bash

find . -name '*.yml' ! -name 'requirements.yml' -d 1 -print0 | xargs -0 "$@"
