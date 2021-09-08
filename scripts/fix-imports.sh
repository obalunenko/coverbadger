#!/bin/sh

set -eu

SCRIPT_NAME="$(basename "$0")"

echo "${SCRIPT_NAME} is running... "

if [ -f "$(go env GOPATH)/bin/gogroup" ] || [ -f "/usr/local/bin/gogroup" ]; then
  gogroup -order std,other,prefix=github.com/obalunenko/ -rewrite $(find . -type f -name "*.go" | grep -v "vendor/" | grep -v ".git")
else
  printf "Cannot check gogroup, please run:
    make install-tools \n"
  exit 1
fi

echo "${SCRIPT_NAME} done."
