#!/usr/bin/env bash

# generate-groups.sh  -> https://github.com/kubernetes/code-generator/blob/master/generate-groups.sh

# work dir
export WORK_DIR=$(cd `dirname $0`; cd ../..; pwd)

# install kubernetes code generator
# go get -u k8s.io/code-generator/...

# generate deepcopy, client, informer, lister by kubernetes code generator
bash ${WORK_DIR}/hack/codegen/generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/sky-big/kubernetes-crd-controller/pkg/client github.com/sky-big/kubernetes-crd-controller/pkg/apis \
  "example:v1alpha1,v1" \
  --go-header-file ${WORK_DIR}/hack/codegen/boilerplate.go.txt

# generate crd resource injection code by knative code generator
bash ${WORK_DIR}/pkg/common/codegen/cmd/injection-gen/generate-injection.sh "injection" \
  github.com/sky-big/kubernetes-crd-controller/pkg/client github.com/sky-big/kubernetes-crd-controller/pkg/apis \
  "example:v1alpha1,v1" \
  --go-header-file ${WORK_DIR}/hack/codegen/boilerplate.go.txt

# generate kubernetes resource injection code by knative code generator
OUTPUT_PKG="github.com/sky-big/kubernetes-crd-controller/pkg/client/kube/injection" \
VERSIONED_CLIENTSET_PKG="k8s.io/client-go/kubernetes" \
EXTERNAL_INFORMER_PKG="k8s.io/client-go/informers" \
bash ${WORK_DIR}/pkg/common/codegen/cmd/injection-gen/generate-injection.sh "injection" \
    k8s.io/client-go \
    k8s.io/api \
    "admissionregistration:v1beta1 apps:v1 autoscaling:v1,v2beta1 batch:v1,v1beta1 core:v1 rbac:v1" \
  --go-header-file ${WORK_DIR}/hack/codegen/boilerplate.go.txt

