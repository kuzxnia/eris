package integration

import (
	"bytes"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/kuzxnia/eris/control-plane/pkg/exception"
	"github.com/kuzxnia/eris/control-plane/tests/integration/helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating workflow", Label("workflow"), func() {
	When("there isn't workflow with same name", func() {
		BeforeEach(func(ctx SpecContext) {
			_, err := WorkflowRepository.GetWorkflow("test one node failover")

			Expect(err).To(MatchError(exception.NotFoundError{}))
		})
		It("creates workflow", func() {
			payload := helpers.ReadFile("data/web/workflow_v1_payload.json")
			request := httptest.NewRequest("POST", "/api/v1/workflow", bytes.NewReader(payload))
			request.Header.Add("Content-Type", "application/json")

			response, _ := FiberApp.Test(request, 1)

			Expect(response).To(HaveHTTPStatus(fiber.StatusCreated))

			workflow, err := WorkflowRepository.GetWorkflow("test one node failover")
			Expect(workflow).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
	})
})
