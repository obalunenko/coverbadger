#!/bin/sh

set -eu

SCRIPT_NAME="$(basename "$0")"

echo "${SCRIPT_NAME} is running... "

GO_FILES=$(find . -type f -name '*.go' | grep -v 'vendor' | grep -v '.git')

echo "gofmt"
gofmt -s -w -l ${GO_FILES}

echo "${SCRIPT_NAME} done."
