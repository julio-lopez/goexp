#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o xtrace


create_key() {
    local -r key_file="$1"
    shift

    echo "Generating SSH key and saving it as '${key_file}'"
    ssh-keygen -q -f "${key_file}" -C "host key" -N '' "$@"
    ssh-keygen -l -f "${key_file}.pub"
}

readonly KEYS_DIR=${1:-"keys"}

mkdir --parents --mode=700 "${KEYS_DIR}"

create_key "${KEYS_DIR}/ssh_host_ed25519_key" -t ed25519
create_key "${KEYS_DIR}/ssh_host_ecdsa_key" -t ecdsa
create_key "${KEYS_DIR}/ssh_host_rsa_key" -t rsa

# Alternatively:
#
# mkdir --parents --mode=700 "${KEYS_DIR}/etc/ssh"
# ssh-keygen -A -f "${KEYS_DIR}"
