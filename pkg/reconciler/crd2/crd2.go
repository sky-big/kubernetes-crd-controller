package crd2

import (
	"context"
)

const (
	// ReconcilerName is the name of the reconciler
	ReconcilerName = "CRD2"
)

type Reconciler struct {
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	return nil
}