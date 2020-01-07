package sharemain

import (
	"flag"
	"log"
	"os"

	"github.com/sky-big/kubernetes-crd-controller/pkg/common/controller"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/kubeconfig"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/logging"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/signals"
	"github.com/sky-big/kubernetes-crd-controller/pkg/common/injection"

	"golang.org/x/sync/errgroup"
	"k8s.io/client-go/rest"
)

func Main(ctors ...controller.ControllerConstructor) {
	var (
		masterURL = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
		config    = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	)
	flag.Parse()

	cfg, err := kubeconfig.GetConfig(*masterURL, *config)
	if err != nil {
		log.Fatal("Error building kubeconfig", err)
	}
	MainWithConfig(cfg, ctors...)
}

func MainWithConfig(cfg *rest.Config, ctors ...controller.ControllerConstructor) {
	// context
	ctx := signals.NewContext()

	// logging
	logger := logging.FromContext(ctx)

	// injection
	ctx, informers := injection.Default.SetupInformers(ctx, cfg)

	// controllers
	controllers := make([]*controller.Impl, 0, len(ctors))
	for _, cf := range ctors {
		ctrl := cf(ctx)
		controllers = append(controllers, ctrl)
	}

	// start informers
	logger.Info("Starting informers.")
	if err := controller.StartInformers(ctx.Done(), informers...); err != nil {
		logger.Fatalw("Failed to start informers", err)
	}

	// start controllers
	logger.Info("Starting controllers.")
	go controller.StartAll(ctx.Done(), controllers...)

	// wait
	eg, egCtx := errgroup.WithContext(ctx)
	<-egCtx.Done()

	if err := eg.Wait(); err != nil {
		logger.Error(err)
		os.Exit(0)
	}
}
