package clients

import (
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/kuzxnia/eris/control-plane/pkg"
	"github.com/kuzxnia/eris/control-plane/pkg/config"
)

type ResourceClient interface {
	GetApiV1Resource(url string, selector *pkg.Selector) ([]*pkg.Resource, error)
}

type HttpResourceClient struct {
	httpClient *resty.Client
}

func ProvideHttpResourceClient(httpClient *resty.Client) *HttpResourceClient {
	return &HttpResourceClient{
		httpClient: httpClient,
	}
}

func (c *HttpResourceClient) GetApiV1Resource(baseUrl string, selector *pkg.Selector) ([]*pkg.Resource, error) {
	url, err := url.JoinPath(baseUrl, config.API_V1_RESOURCE_ENDPOINT)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.R().
		SetResult([]*pkg.Resource{}).
		SetQueryParams(map[string]string{}).
		Get(url)
	if err != nil {
		return nil, err
	}
	return *(resp.Result().(*[]*pkg.Resource)), nil
}
