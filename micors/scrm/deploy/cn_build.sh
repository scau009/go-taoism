#!/usr/bin/env bash

basepath=$(cd `dirname $0`; pwd);

go env -w  GOPROXY=https://goproxy.cn,direct

cd ${basepath}/../bin/http/overseer;

go build -ldflags "-X main.BuildID=$(date +%s)" -o main_latest main.go;
