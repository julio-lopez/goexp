#!/usr/bin/env bash

set -o errexit
set -o nounset
set +o xtrace

pushd /mnt

git clone https://github.com/kastenhq/kopia
