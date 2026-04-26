#!/usr/bin/env sh

set -o errexit
set -o nounset
set -o xtrace

readonly s6_svc_dir="/home/s6/active-svc"

cd "${s6_svc_dir}"
ln -s ../services/ssh .

s6-svscanctl -h "${s6_svc_dir}"
