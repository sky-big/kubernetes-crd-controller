package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sky-big/kubernetes-crd-controller/pkg/common/controller"
	"github.com/sky-big/kubernetes-crd-controller/pkg/reconciler/crd1"
	"github.com/sky-big/kubernetes-crd-controller/pkg/reconciler/crd2"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/signals"

	"golang.org/x/sync/errgroup"
)

type ControllerConstructor func(context.Context) *controller.Impl

func main() {
	ctx := signals.NewContext()
	ctors := []ControllerConstructor{
		crd1.NewController,
		crd2.NewController,
	}
	controllers := make([]*controller.Impl, 0, len(ctors))
	for _, cf := range ctors {
		ctrl := cf(ctx)
		controllers = append(controllers, ctrl)
	}

	go controller.StartAll(ctx.Done(), controllers...)

	eg, egCtx := errgroup.WithContext(ctx)
	<-egCtx.Done()

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}