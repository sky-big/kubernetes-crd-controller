package crd2

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
	ReconcilerName = "CRD2"
)

type Reconciler struct {
	crd2Informer v1alpha1.CRD2Informer
	crd2Lister   listers.CRD2Lister
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	logger := logging.FromContext(ctx)

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		logger.Errorw("Invalid resource key", zap.Error(err))
		return nil
	}

	// Get crd1 resource with this namespace/name
	original, err := c.crd2Lister.CRD2s(namespace).Get(name)

	logger.Infof("Reconcile V1Alpha1 CRD2 Resource %+v", original)
	return nil
}
