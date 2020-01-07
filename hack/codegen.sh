#!/usr/bin/env bash

# generate-groups.sh  -> https://github.com/kubernetes/code-generator/blob/master/generate-groups.sh
# generate-knative.sh -> https://github.com/knative/pkg/blob/master/hack/generate-knative.sh

# work dir
export WORK_DIR=$(cd `dirname $0`; cd ..; pwd)

# install kubernetes code generator
go get -u k8s.io/code-generator/...

# install knative injection code generator
go get -u knative.dev/pkg/codegen/...

# generate deepcopy, client, informer, lister by kubernetes code generator
bash ${WORK_DIR}/hack/generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/sky-big/kubernetes-crd-controller/pkg/client github.com/sky-big/kubernetes-crd-controller/pkg/apis \
  "example:v1alpha1,v1" \
  --go-header-file ${WORK_DIR}/hack/boilerplate.go.txt

# generate injection code by knative code generator
bash ${WORK_DIR}/hack/generate-knative.sh "injection" \
  github.com/sky-big/kubernetes-crd-controller/pkg/client github.com/sky-big/kubernetes-crd-controller/pkg/apis \
  "example:v1alpha1,v1" \
  --go-header-file ${WORK_DIR}/hack/boilerplate.go.txt
