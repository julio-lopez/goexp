#!/usr/bin/env sh

set -o errexit
set -o nounset
set -o xtrace


cd "${HOME}/service"
ln -s ../services/ssh .

s6-svscanctl -h "${HOME}/service"
