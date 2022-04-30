#!/usr/bin/env bash

set -o errexit
set -o nounset
set +o xtrace

readonly SRC_DIR=${PATH_PREFIX:-/${HOME}}/go/src

mkdir -p ${SRC_DIR}

pushd ${SRC_DIR}

git clone https://github.com/kopia/kopia

cd kopia

go build .

# make build-current-os-noui
