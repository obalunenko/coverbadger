#!/bin/sh

set -eu

SCRIPT_NAME="$(basename "$0")"

echo "${SCRIPT_NAME} is running... "

go generate -x ./...

echo "${SCRIPT_NAME} done."
