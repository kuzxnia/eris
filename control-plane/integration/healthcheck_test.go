package integration

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	withApp(t, func(app *fiber.App) {
		// given
		req := httptest.NewRequest("GET", "/api/v1/health", nil)

		// when
		resp, _ := app.Test(req, 1)

		// then
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
