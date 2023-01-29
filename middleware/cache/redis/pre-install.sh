#!/usr/bin/env bash
set -eu
server="207.148.113.13"
redisDir="/data/redis/data"
if [[ ! -d ${redisDir} ]]; then
    mkdir -pv ${redisDir}
fi
echo "done"

