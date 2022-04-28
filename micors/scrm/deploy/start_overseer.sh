#!/usr/bin/env bash

basepath=$(cd `dirname $0`; pwd);

cd ${basepath}/../bin/http/overseer;

./main_latest;
