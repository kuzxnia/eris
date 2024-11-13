package mocks

import (
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

func ProvideHttpClientMock() *resty.Client {
	client := resty.New()

	// Get the underlying HTTP Client and set it to Mock
	httpmock.ActivateNonDefault(client.GetClient())

	return client
}
