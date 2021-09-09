#!/usr/bin/env bash

set -Eeuo pipefail

function cleanup() {
  trap - SIGINT SIGTERM ERR EXIT
  echo "cleanup running"
}

trap cleanup SIGINT SIGTERM ERR EXIT

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"

echo "${SCRIPT_NAME} is running... "

# Get new tags from the remote
git fetch --tags -f

export BUILD_COMMIT=$(git rev-parse HEAD)
export BUILD_SHORTCOMMIT=$(git rev-parse --short HEAD)
export BUILD_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)
export BUILD_VERSION=$(git describe --tags  --always $(git rev-list --tags --max-count=1))

export BUILD_APPNAME=coverbadger

if [ -z "${BUILD_VERSION}" ] || [ "${BUILD_VERSION}" = "${BUILD_SHORTCOMMIT}" ]
 then
  BUILD_VERSION="v0.0.0"
fi

export BUILD_GOVERSION=$(go version | awk '{print $3;}')

goreleaser --snapshot --skip-publish --rm-dist