package integration

import (
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthCheck", func() {
	var request *http.Request

	BeforeEach(func() {
		request = httptest.NewRequest("GET", "/api/v1/health", nil)
	})

	It("send request", func() {
		response, _ := FiberApp.Test(request, 1)
		Expect(response.StatusCode).To(Equal(fiber.StatusOK))
	})
})
