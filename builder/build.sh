#!/usr/bin/env bash
set -e
docker build --no-cache --platform=linux/amd64 -t saichler/netop-builder:latest .
docker push saichler/netop-builder:latest
