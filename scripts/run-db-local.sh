#!/bin/sh

docker run -d \
  --name nb-psql \
  -p 25432:5432 \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=nb \
  postgres:14-alpine
