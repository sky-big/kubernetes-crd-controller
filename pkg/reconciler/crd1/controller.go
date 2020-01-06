package crd1

import (
	"context"

	"github.com/sky-big/kubernetes-crd-controller/pkg/common/controller"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/logging"
)

func NewController(
	ctx context.Context,
) *controller.Impl {
	logger := logging.FromContext(ctx)

	c := &Reconciler{}
	impl := controller.NewImpl(c, logger, ReconcilerName)

	logger.Info("CRD1 Controller Started")
	return impl
}
