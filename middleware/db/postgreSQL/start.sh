#!/usr/bin/env bash

docker run \
    -d \
    -e POSTGRES_HOST_AUTH_METHOD=trust \
    -e POSTGRES_USER=jun \
    -e POSTGRES_PASSWORD=Initial0! \
    -e POSTGRES_DB=shoe \
    -p 5432:5432 \
    postgres:14.4

