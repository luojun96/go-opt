#!/usr/bin/env bash

tag="7.0.4"
docker run \
    -d \
    -p 6379:6379 \
    -v /data/redis/redis.conf:/etc/redis/redis.conf \
    -v /data/redis/data:/data \
    redis:${tag} \
    redis-server /etc/redis/redis.conf
