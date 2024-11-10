package workflow_test

import (
	"github.com/kuzxnia/eris/control-plane/internal/infra/repository"
	"github.com/kuzxnia/eris/control-plane/pkg/exception"
	"github.com/kuzxnia/eris/control-plane/pkg/interfaces/web/dto"
	"github.com/kuzxnia/eris/control-plane/pkg/workflow"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating workflow", Label("workflow"), func() {
	var workflowService *workflow.WorkflowService
	var workflowRepository workflow.WorkflowRepository

	BeforeEach(func(ctx SpecContext) {
		workflowRepository = &repository.InMemoryWorkflowRepository{}
		workflowService = workflow.ProvideWorkflowService(nil, &workflow.WorkflowMapper{}, workflowRepository)
	})

	When("there isn't workflow with same name", func() {
		BeforeEach(func(ctx SpecContext) {
			_, err := workflowRepository.GetWorkflow("test-workflow")

			Expect(err).To(MatchError(exception.NotFoundError{}))
		})

		It("creates workflow", func() {
			dummyWorkflow := dto.WorkflowRequest{
				Name: "test-workflow",
			}

			Expect(workflowService.CreateWorkflow(dummyWorkflow)).To(Succeed())

			workflow, err := workflowRepository.GetWorkflow(dummyWorkflow.Name)
			Expect(workflow).To(Equal(dummyWorkflow))
			Expect(err).To(BeNil())
		})
	})

	When("there is workflow with same name", func() {
		BeforeEach(func(ctx SpecContext) {
			Expect(
				workflowRepository.SaveWorkflow(&workflow.Workflow{Name: "test two node failover"}),
			).To(Succeed())
		})
		It("fails to create workflow", func() {
			// payload := helpers.ReadFile("data/web/another_workflow_v1_payload.json")
			// request := httptest.NewRequest("POST", "/api/v1/workflow", bytes.NewReader(payload))
			// request.Header.Add("Content-Type", "application/json")

			// response, _ := FiberApp.Test(request, 1)

			// Expect(response.StatusCode).To(Equal(fiber.StatusConflict))
		})
	})
})
