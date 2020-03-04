#!/usr/bin/env sh
set -o nounset
set -o errexit
set -o xtrace


# env | sort
echo "\$0: '${0}'"
echo "PWD='${PWD}'"
pwd
cd hello
go build
