#!/bin/bash

export WORK_DIR=$(cd `dirname $0`; pwd)
cd ${WORK_DIR}

IMAGE=skybig/kubernetes-crd-controller:latest

# build controller
cd .. && make build && cd ./docker

# get controller bin
cp ../bin/kubernetes-crd-controller .

echo "[START] build controller images"

# build docker image
docker build --tag "${IMAGE}" .

echo "[END] build controller images"

# remove controller bin
rm -f ./kubernetes-crd-controller
