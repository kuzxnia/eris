package integration

import (
	"bytes"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/kuzxnia/eris/control-plane/integration/helpers"
)

func (s *TestSuite) TestAgent() {
	// given
	payload := helpers.ReadFile("data/web/agent_v1_payload.json")
	request := httptest.NewRequest("POST", "/api/v1/agent", bytes.NewReader(payload))
	request.Header.Add("Content-Type", "application/json")

	// when
	resp, _ := s.FiberApp.Test(request, 1)

	// then
	s.Equal(fiber.StatusCreated, resp.StatusCode)
}
