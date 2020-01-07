package main

import (
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/sharemain"
	"github.com/sky-big/kubernetes-crd-controller/pkg/reconciler/crd1"
	"github.com/sky-big/kubernetes-crd-controller/pkg/reconciler/crd2"
)

func main() {
	sharemain.Main(
		crd1.NewController,
		crd2.NewController,
	)
}
