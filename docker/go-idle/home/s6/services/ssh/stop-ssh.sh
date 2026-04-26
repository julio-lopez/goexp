#!/usr/bin/env sh

set -o errexit
set -o nounset
set -o xtrace

readonly s6_svc_dir="/home/s6/active-svc"
 
rm -f "${s6_svc_dir}/ssh"

s6-svscanctl -h "${s6_svc_dir}"
