#!/usr/bin/env bash

set -euo pipefail

just -f ./playbooks/justfile playbook --ask-become-pass "$@" main.yml

git status
yadm status
