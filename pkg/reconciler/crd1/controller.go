package crd1

import (
	"context"

	"github.com/sky-big/kubernetes-crd-controller/pkg/client/injection/informers/example/v1alpha1/crd1"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/controller"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/logging"
)

func NewController(
	ctx context.Context,
) *controller.Impl {
	logger := logging.FromContext(ctx)
	crd1Informer := crd1.Get(ctx)

	c := &Reconciler{
		crd1Informer: crd1Informer,
		crd1Lister:   crd1Informer.Lister(),
	}
	impl := controller.NewImpl(c, logger, ReconcilerName)

	crd1Informer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	logger.Info("CRD1 Controller Started")
	return impl
}
