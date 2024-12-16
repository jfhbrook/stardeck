#!/usr/bin/env bash

if [[ "${1}" == enable ]]; then
  pactl load-module module-loopback latency_msec=1
elif [[ "${1}" == disable ]]; then
  pactl unload-module module-loopback
else
  echo "Unknown command: ${1}"
  exit 1
fi
