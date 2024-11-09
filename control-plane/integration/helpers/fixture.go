package helpers

import (
	"net/http"
	"os"
	"testing"

	main "github.com/kuzxnia/eris/control-plane/pkg"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func withApp(t *testing.T, r interface{}) {
	app := fxtest.New(t,
		main.Modules,
		// fx.Replace(mocks(t)),
		fx.Invoke(r),
	)
	defer app.RequireStop()
	app.RequireStart()
}

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func ResponseBodyAsString(response *http.Response) string {
	bodyData := make([]byte, response.ContentLength)
	_, _ = response.Body.Read(bodyData)
	return string(bodyData)
}