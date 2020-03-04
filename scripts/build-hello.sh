#!/usr/bin/env sh
set -o nounset
set -o errexit
set -o xtrace


# env | sort
echo "\$0: '${0}'"
echo "PWD='${PWD}'"
pwd
go build -o hello/hello ./hello/
