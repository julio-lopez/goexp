#!/usr/bin/env bash

set -o errexit
set -o nounset
set +o xtrace

readonly SRC_DIR=${PATH_PREFIX:-/mnt}/src

mkdir -p ${SRC_DIR}

pushd ${SRC_DIR}

git clone https://github.com/kastenhq/kopia

cd kopia

make build-current-os-noui
