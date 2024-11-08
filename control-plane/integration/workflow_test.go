package integration

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestWokflowCreate(t *testing.T) {
	withApp(t, func(app *fiber.App) {
		// given
		payload := readFile("data/web/workflow_v1_payload.json")
		req := httptest.NewRequest("POST", "/api/v1/workflow", bytes.NewReader(payload))
		req.Header.Add("Content-Type", "application/json")

		// when
		resp, _ := app.Test(req)

		// then
    // data := responseBodyAsString(resp)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
