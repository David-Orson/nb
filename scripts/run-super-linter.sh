#!/bin/sh

ROOT_DIR="$(pwd)/$(dirname "$0")/.."

docker run \
    -e RUN_LOCAL=true \
    -e USE_FIND_ALGORITHM=true \
    -v "$ROOT_DIR":/tmp/lint \
    nb-super-linter


