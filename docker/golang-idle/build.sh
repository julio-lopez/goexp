#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o xtrace
#set -o pipefail

readonly ts="$(date +%Y.%m.%d-%H%M%S)"
readonly basedir=$(realpath --logical --canonicalize-existing "$(dirname ${0})")

cd "${basedir}"

docker build \
    --build-arg imageVersion="${ts}" \
    --tag "kanisterio/golang:bare-${ts}" .

cd -
