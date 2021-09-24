#!/bin/sh

set -eu

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd ${SCRIPT_DIR} && git rev-parse --show-toplevel)"
BIN_DIR=${REPO_ROOT}/bin
COV_DIR=${REPO_ROOT}/coverage

echo "${SCRIPT_NAME} is running... "

APP=coverbadger

COVERAGE=$(gocov report ${COV_DIR}/full.json | tail -1 | awk '{if ($1 != "?") print $3; else print "0.0";}' | sed 's/\%//g')

${BIN_DIR}/${APP} \
  --coverage=${COVERAGE} \
  --md="README.md" \
  -style=flat
