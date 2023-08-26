#!/usr/bin/env sh
set -o nounset
set -o errexit
set -o xtrace

base=$(dirname ${0})
TARGET_DIR=$(realpath --logical --canonicalize-existing "${base}/../hello")

docker build \
    --tag juliolopez/hello \
    "${TARGET_DIR}"
