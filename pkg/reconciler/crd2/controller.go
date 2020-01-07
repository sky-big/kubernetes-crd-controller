package crd2

import (
	"context"

	"github.com/sky-big/kubernetes-crd-controller/pkg/client/injection/informers/example/v1alpha1/crd2"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/controller"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/logging"
)

// NewController initializes the controller and is called by the generated code
// Registers eventhandlers to enqueue events
func NewController(
	ctx context.Context,
) *controller.Impl {
	logger := logging.FromContext(ctx)
	crd2Informer := crd2.Get(ctx)

	c := &Reconciler{
		crd2V1alpha1Informer: crd2Informer,
		crd2V1alpha1Lister:   crd2Informer.Lister(),
	}
	impl := controller.NewImpl(c, logger, ReconcilerName)

	logger.Info("CRD2 Controller Started")
	return impl
}
