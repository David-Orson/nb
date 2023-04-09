#!/bin/sh

PROJECT_DIR="/go/src/api"

cd "$(dirname "$0")/../api/" || exit 1

docker run -it --rm \
    -w "${PROJECT_DIR}" \
    -e "NB_DB_HOST=localhost" \
    -e "NB_DB_PORT=25432" \
    -e "NB_DB_USER=postgres" \
    -e "NB_DB_PASSWORD=password" \
    -e "NB_DB_NAME=nb" \
    -e "air_wd=${PROJECT_DIR}" \
    -v "$(pwd)"/:"${PROJECT_DIR}" \
    -v "$GOPATH"/pkg/mod:/go/pkg/mod \
    --network=host \
    cosmtrek/air:v1.40.4
