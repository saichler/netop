#!/usr/bin/env bash
set -e
docker build --no-cache --platform=linux/amd64 -t saichler/netop-device-inv:latest .
docker push saichler/netop-device-inv:latest
