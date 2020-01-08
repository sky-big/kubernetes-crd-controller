package crd1

import (
	"context"

	"github.com/sky-big/kubernetes-crd-controller/pkg/client/informers/externalversions/example/v1alpha1"
	listers "github.com/sky-big/kubernetes-crd-controller/pkg/client/listers/example/v1alpha1"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/logging"

	"go.uber.org/zap"
	"k8s.io/client-go/tools/cache"
)

const (
	// ReconcilerName is the name of the reconciler
	ReconcilerName = "CRD1"
)

type Reconciler struct {
	crd1Informer v1alpha1.CRD1Informer
	crd1Lister   listers.CRD1Lister
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	logger := logging.FromContext(ctx)

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		logger.Errorw("Invalid resource key", zap.Error(err))
		return nil
	}

	// Get crd1 resource with this namespace/name
	original, err := c.crd1Lister.CRD1s(namespace).Get(name)

	logger.Infof("Reconcile CRD1 Resource %+v", original)
	return nil
}
