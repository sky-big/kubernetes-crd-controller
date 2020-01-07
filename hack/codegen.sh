#!/usr/bin/env bash

# generate-groups.sh  -> https://github.com/kubernetes/code-generator/blob/master/generate-groups.sh

# work dir
export WORK_DIR=$(cd `dirname $0`; cd ..; pwd)

# install kubernetes code generator
go get -u k8s.io/code-generator/...

# generate deepcopy, client, informer, lister by kubernetes code generator
bash ${WORK_DIR}/hack/generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/sky-big/kubernetes-crd-controller/pkg/client github.com/sky-big/kubernetes-crd-controller/pkg/apis \
  "example:v1alpha1,v1" \
  --go-header-file ${WORK_DIR}/hack/boilerplate.go.txt

# generate injection code by knative code generator
bash ${WORK_DIR}/pkg/common/codegen/cmd/injection-gen/generate-injection.sh "injection" \
  github.com/sky-big/kubernetes-crd-controller/pkg/client github.com/sky-big/kubernetes-crd-controller/pkg/apis \
  "example:v1alpha1,v1" \
  --go-header-file ${WORK_DIR}/hack/boilerplate.go.txt
