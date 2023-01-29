#!/usr/bin/env bash
set -eu
server="207.148.113.13"
redisConfDir="/data/redis"

cd $(dirname $0)

ssh -o StrictHostKeyChecking=no root@${server} < pre-install.sh
scp -o StrictHostKeyChecking=no redis.conf root@${server}:${redisConfDir}

