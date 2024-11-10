package workflow_test

import (
	"github.com/kuzxnia/eris/control-plane/pkg/exception"
	"github.com/kuzxnia/eris/control-plane/pkg/interfaces/web/dto"
	"github.com/kuzxnia/eris/control-plane/pkg/workflow"
	"github.com/kuzxnia/eris/control-plane/tests/unit/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Creating workflow", Label("workflow"), func() {
	var workflowService *workflow.WorkflowService
	// var workflowRepository *mocks.WorkflowRepositoryMock
	var workflowRepository *mocks.MockWorkflowRepository

	BeforeEach(func(ctx SpecContext) {
		mockCtrl := gomock.NewController(GinkgoT())

		workflowRepository = mocks.NewMockWorkflowRepository(mockCtrl)
		workflowService = workflow.ProvideWorkflowService(
			&workflow.WorkflowMapper{}, workflowRepository,
		)
	})

	When("there isn't workflow with same name", func() {
		BeforeEach(func(ctx SpecContext) {
			workflowRepository.
				EXPECT().
				SaveWorkflow(gomock.Any())
		})

		It("creates workflow", func(ctx SpecContext) {
			dummyWorkflow := dto.WorkflowRequest{
				Name: "test-workflow",
			}

			Expect(workflowService.CreateWorkflow(dummyWorkflow)).To(Succeed())
		})
	})

	When("there is workflow with same name", func() {
		BeforeEach(func(ctx SpecContext) {
			workflowRepository.
				EXPECT().
				SaveWorkflow(gomock.Any()).
				Return(exception.ResourceAlreadyExistsError{})
		})
		It("fails to create workflow", func() {
			dummyWorkflow := dto.WorkflowRequest{
				Name: "test-workflow",
			}

			Expect(workflowService.CreateWorkflow(dummyWorkflow)).Error().Should(HaveOccurred())
		})
	})
})
