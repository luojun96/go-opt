#!/usr/bin/env bash

tag="8.0.29"
docker run \
    -d \
    -e MYSQL_ROOT_PASSWORD=Initial0! \
    -e MYSQL_DATABASE=shoes \
    -p 3306:3306 \
    mysql:${tag}
