#!/usr/bin/env bash

basepath=$(cd `dirname $0`; pwd);

cd ${basepath};

git pull;

docker-compose run --rm app /bin/bash -x ./deploy/cn_build.sh