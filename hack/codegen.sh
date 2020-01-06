#!/usr/bin/env bash

go get -u k8s.io/code-generator/...

cd $GOPATH/src/k8s.io/code-generator

./generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/sky-big/kubernetes-crd-controller/pkg/client github.com/sky-big/kubernetes-crd-controller/pkg/apis \
  "example:v1alpha1,v1" \
