#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"

# shellcheck disable=SC2046
echo "${SCRIPT_NAME} is running... "

# shellcheck disable=SC1090
source "${SCRIPT_DIR}"/linters-source.sh

vet

echo "${SCRIPT_NAME} done."
