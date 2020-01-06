package crd1

import (
	"context"
)

const (
	// ReconcilerName is the name of the reconciler
	ReconcilerName = "CRD1"
)

type Reconciler struct {
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	return nil
}