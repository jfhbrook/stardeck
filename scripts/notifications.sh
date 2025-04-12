#!/usr/bin/env bash

busctl monitor \
  --user \
  --destination=org.freedesktop.Notifications \
  --match="member='Notify'" \
  --json=short
