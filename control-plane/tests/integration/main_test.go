package integration

import (
	"context"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kuzxnia/eris/control-plane/internal/infra"
	"github.com/kuzxnia/eris/control-plane/pkg/agent"
	"github.com/kuzxnia/eris/control-plane/pkg/config"
	"github.com/kuzxnia/eris/control-plane/pkg/interfaces/web"
	"github.com/kuzxnia/eris/control-plane/pkg/workflow"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

var (
	FxApp              *fx.App
	FiberApp           *fiber.App
	WorkflowRepository workflow.WorkflowRepository
)

var _ = SynchronizedBeforeSuite(func() []byte {
	// runs *only* on process #1
	FxApp = fx.New(
		config.Module,
		infra.Module,
		agent.Module,
		workflow.Module,
		web.Module,
		// fx.Replace(mocks(t)),
		// fx.Invoke(r),
		fx.Invoke(func(fiberApp *fiber.App, workflowRepository workflow.WorkflowRepository) {
			FiberApp = fiberApp
			WorkflowRepository = workflowRepository
		}),
	)
	Expect(FxApp.Start(context.Background())).To(Succeed())
	// dbRunner := db.NewRunner()
	// Expect(dbRunner.Start()).To(Succeed())
	return []byte{}
}, func(address []byte) {
	// runs on *all* processes
	// dbClient = db.NewClient()
	// Expect(dbClient.Connect(string(address))).To(Succeed())
	// dbClient.SetNamespace(fmt.Sprintf("namespace-%d", GinkgoParallelProcess()))
})

var _ = SynchronizedAfterSuite(func() {
	// runs on *all* processes
	// Expect(dbClient.Cleanup()).To(Succeed())
}, func() {
	// runs *only* on process #1
	// Expect(dbRunner.Stop()).To(Succeed())
	Expect(FxApp.Stop(context.Background())).To(Succeed())
})

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Suite")
}
