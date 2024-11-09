package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/kuzxnia/eris/control-plane/integration/helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Workflow", func() {
	Describe("Creating workflow", func() {
		var payload []byte
		var request *http.Request

		BeforeEach(func() {
			payload = helpers.ReadFile("data/web/workflow_v1_payload.json")
			request = httptest.NewRequest("POST", "/api/v1/workflow", bytes.NewReader(payload))
			request.Header.Add("Content-Type", "application/json")
		})

		It("send request", func() {
			response, _ := FiberApp.Test(request, 1)
			Expect(response.StatusCode).To(Equal(fiber.StatusCreated))
		})
	})
})
