#!/usr/bin/env bash

# TODO: Node script that reuses config logic
ANSIBLE_CONFIG="${ANSIBLE_CONFIG:-$(pwd)/config/ansible.cfg}"
ANSIBLE_ROLES_PATH="${ANSIBLE_ROLES_PATH:-$(pwd)/roles}"
ANSIBLE_COLLECTIONS_PATH="${ANSIBLE_COLLECTIONS_PATH:-$(pwd)/collections}"

export ANSIBLE_CONFIG
export ANSIBLE_ROLES_PATH
export ANSIBLE_COLLECTIONS_PATH

ansible-galaxy install -r ./requirements.yml "$@"
