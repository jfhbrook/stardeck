#!/usr/bin/env bash

source ../../.copr/bin/prelude.sh

set-gh-release-version 45Drives/cockpit-file-sharing
download-sources
build-srpm
