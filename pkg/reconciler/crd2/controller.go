package crd2

import (
	"context"

	"github.com/sky-big/kubernetes-crd-controller/pkg/common/controller"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/logging"
)

// NewController initializes the controller and is called by the generated code
// Registers eventhandlers to enqueue events
func NewController(
	ctx context.Context,
) *controller.Impl {

	c := &Reconciler{}
	impl := controller.NewImpl(c, logging.FromContext(ctx), ReconcilerName)

	return impl
}
