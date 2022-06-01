#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o xtrace
#set -o pipefail

readonly IMAGE="ghcr.io/julio-lopez/go-idle"
readonly buildVersion="$(git log -1 --format=%ad-%h --date=format:'%Y.%m.%d-%H%m')"
readonly ts="$(date +%Y-%m-%d-%H%M%S)"
readonly basedir=$(realpath --logical --canonicalize-existing "$(dirname ${0})")

cd "${basedir}"


docker build \
    --build-arg imageVersion="${buildVersion}" \
    --tag "${IMAGE}:${buildVersion}" \
    --tag "${IMAGE}:${buildVersion}-${ts}" \
    --tag "${IMAGE}:latest" \
    .

cd -
