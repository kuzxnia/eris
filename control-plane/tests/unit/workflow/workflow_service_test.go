package workflow_test

import (
	"testing"

	"github.com/kuzxnia/eris/control-plane/pkg/exception"
	"github.com/kuzxnia/eris/control-plane/pkg/resource"
	"github.com/kuzxnia/eris/control-plane/pkg/web/dto"
	"github.com/kuzxnia/eris/control-plane/pkg/workflow"
	"github.com/kuzxnia/eris/control-plane/tests/unit/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Configuring workflow", Label("workflow"), func() {
	var workflowService *workflow.WorkflowService
	var workflowRepository *mocks.MockWorkflowRepository
	var workflowConfigRepository *mocks.MockWorkflowConfigRepository

	BeforeEach(func(ctx SpecContext) {
		mockCtrl := gomock.NewController(GinkgoT())

		workflowRepository = mocks.NewMockWorkflowRepository(mockCtrl)
		workflowConfigRepository = mocks.NewMockWorkflowConfigRepository(mockCtrl)
		workflowService = workflow.ProvideWorkflowService(
			&workflow.WorkflowMapper{}, workflowRepository, workflowConfigRepository, nil, nil,
		)
	})

	When("there isn't config with same name", func() {
		BeforeEach(func(ctx SpecContext) {
			workflowConfigRepository.
				EXPECT().
				CreateWorkflowConfig(gomock.Any())
		})

		It("creates config", func(ctx SpecContext) {
			dummyWorkflow := dto.WorkflowConfigRequest{
				Name: "test-workflow",
			}

			Expect(workflowService.ConfigureWorkflow(dummyWorkflow)).To(Succeed())
		})
	})

	When("there isn't config with same name", func() {
		BeforeEach(func(ctx SpecContext) {
			workflowConfigRepository.
				EXPECT().
				CreateWorkflowConfig(gomock.Any()).
				Return(exception.ResourceAlreadyExistsError{})
		})

		It("raises resource already exists error", func(ctx SpecContext) {
			dummyWorkflow := dto.WorkflowConfigRequest{
				Name: "test-workflow",
			}

			Expect(workflowService.ConfigureWorkflow(dummyWorkflow)).
				Error().To(Equal(exception.ResourceAlreadyExistsError{}))
		})
	})
})

var _ = Describe("Running workflow", Label("workflow"), func() {
	var workflowService *workflow.WorkflowService
	var workflowRepository *mocks.MockWorkflowRepository
	var workflowConfigRepository *mocks.MockWorkflowConfigRepository
	var actionExecutor *mocks.MockActionExecutor
	var defaultResourceSelector *mocks.MockResourceSelector

	BeforeEach(func(ctx SpecContext) {
		mockCtrl := gomock.NewController(GinkgoT())

		workflowRepository = mocks.NewMockWorkflowRepository(mockCtrl)
		workflowConfigRepository = mocks.NewMockWorkflowConfigRepository(mockCtrl)
		workflowConfigRepository = mocks.NewMockWorkflowConfigRepository(mockCtrl)
		actionExecutor = mocks.NewMockActionExecutor(mockCtrl)
		defaultResourceSelector = mocks.NewMockResourceSelector(mockCtrl)
		defaultResourceSelector.EXPECT().GetSelectorType().AnyTimes().Return("dummy-selector-type")
		selectorRegistry := resource.ProvideResourceSelectorRegistry(
			[]resource.ResourceSelector{
				defaultResourceSelector,
			},
		)

		workflowService = workflow.ProvideWorkflowService(
			&workflow.WorkflowMapper{}, workflowRepository, workflowConfigRepository, selectorRegistry, actionExecutor,
		)
	})

	When("workflow is configured", func() {
		BeforeEach(func(ctx SpecContext) {
			dummyWorkflowConfig := workflow.WorkflowConfig{
				Name: "dummy-name",
				Actions: []*workflow.Action{
					{
						Type:     "dummy-selector-type",
						Selector: &workflow.Selector{},
					},
				},
			}

			workflowConfigRepository.
				EXPECT().GetWorkflowConfig("dummy-name").Return(&dummyWorkflowConfig, nil)

			workflowRepository.
				EXPECT().
				CreateWorkflow(
					gomock.Cond(func(x *workflow.Workflow) bool {
						return x.Config == &dummyWorkflowConfig
					}),
				)
		})
		It("creates workflow and saved to storage", func() {
			Expect(workflowService.CreateWorkflowFromConfig("dummy-name")).To(Succeed())
		})

		It("resolves resources for actions", func() {
		})

		It("fails to resolve resources for actions", func() {
		})

		It("saves resolved resources for actions", func() {
		})

		It("sends action for execution", func() {
		})
	})

	When("workflow isn't configured", func() {
		BeforeEach(func(ctx SpecContext) {
			workflowConfigRepository.
				EXPECT().GetWorkflowConfig("dummy-name").Return(nil, exception.NotFoundError{})
		})

		It("fails to create workflow", func(ctx SpecContext) {
			Expect(workflowService.StartWorkflow("dummy-name")).
				Error().To(Equal(exception.NotFoundError{}))
		})
	})
})

func TestWorkflow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WorkflowServiceSuite")
}
