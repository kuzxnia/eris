package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/kuzxnia/eris/control-plane/tests/integration/helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Agent", func() {
	Describe("Creating agent", func() {
		When("there isn't agent with same Id", func() {
			var payload []byte
			var request *http.Request

			BeforeEach(func() {
				payload = helpers.ReadFile("data/web/agent_v1_payload.json")
				request = httptest.NewRequest("POST", "/api/v1/agent", bytes.NewReader(payload))
				request.Header.Add("Content-Type", "application/json")
			})

			It("creates agent", func() {
				response, _ := FiberApp.Test(request, 1)
				Expect(response.StatusCode).To(Equal(fiber.StatusCreated))
			})
		})
	})
})
