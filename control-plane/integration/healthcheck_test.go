package integration

import (
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

func (s *TestSuite) TestHealthCheck() {
	// given
	req := httptest.NewRequest("GET", "/api/v1/health", nil)

	// when
	resp, _ := s.FiberApp.Test(req, 1)

	// then
	s.Equal(fiber.StatusOK, resp.StatusCode)
}
