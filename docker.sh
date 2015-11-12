#!/usr/bin/env bash

set -e

docker run \
  --name GoDash \
  --rm \
  -ti \
  -v $(pwd):/root/GoDash/src/godash \
  -e GOPATH=/go:/root/GoDash \
  mooxe/golang \
  /bin/bash
