package helpers

import (
	"net/http"
	"os"
)

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
