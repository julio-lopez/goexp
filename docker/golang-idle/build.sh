#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o xtrace
#set -o pipefail

readonly buildVersion="$(git log -1 --format=%ad-%h --date=format:'%Y.%d.%m-%H%m')"
readonly ts="$(date +%Y.%m.%d-%H%M%S)"
readonly basedir=$(realpath --logical --canonicalize-existing "$(dirname ${0})")

cd "${basedir}"

docker build \
    --build-arg imageVersion="${buildVersion}" \
    --label buildVersion="${buildVersion}" \
    --label buildDate="${ts}" \
    --tag "ghcr.io/julio-lopez/golang-idle:${buildVersion}" .

cd -
