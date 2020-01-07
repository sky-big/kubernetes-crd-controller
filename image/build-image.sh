#!/bin/bash

export WORK_DIR=$(cd `dirname $0`; pwd)
cd ${WORK_DIR}

IMAGE=skybig/kubernetes-crd-controller:latest

# build controller
cd ${WORK_DIR}/.. && make build && cd ${WORK_DIR}

# get controller bin
cp ${WORK_DIR}/../bin/kubernetes-crd-controller .

echo "[START] build controller images"

# build docker image
docker build --tag "${IMAGE}" .

echo "[END] build controller images"

# remove controller bin
rm -f ./kubernetes-crd-controller
