#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o xtrace
#set -o pipefail

readonly IMAGE="ghcr.io/julio-lopez/idle"
readonly buildVersion="$(git log -1 --format=%ad-%h --date=format:'%Y.%m.%d-%H%m%S')"
readonly ts="$(date +%Y-%m-%d-%H%M%S)"
readonly basedir=$(realpath --logical --canonicalize-existing "$(dirname ${0})")

cd "${basedir}"

#     --tag "${IMAGE}:${buildVersion}-${ts}" \

docker build \
    --build-arg imageVersion="${buildVersion}" \
    --tag "${IMAGE}:${buildVersion}" \
    --tag "${IMAGE}:latest" \
    .

cd -
