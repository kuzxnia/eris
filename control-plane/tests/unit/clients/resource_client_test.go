package clients_test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/kuzxnia/eris/control-plane/pkg"
	"github.com/kuzxnia/eris/control-plane/pkg/clients"
	"github.com/kuzxnia/eris/control-plane/tests/unit/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getting resources external service", Ordered, Label("clients"), func() {
	var httpClient *resty.Client
	var resourceClient clients.ResourceClient
	BeforeAll(func() {
		httpClient = mocks.ProvideHttpClientMock()
	})

	BeforeEach(func(ctx SpecContext) {
		resourceClient = clients.ProvideHttpResourceClient(httpClient)
	})
	AfterEach(func(ctx SpecContext) {
		httpmock.DeactivateAndReset()
	})

	When("provided url is correct", func() {
		baseUrl := "http://example-base-url.com"

		BeforeEach(func(ctx SpecContext) {
			responder, _ := httpmock.NewJsonResponder(200, []map[string]string{
				{"name": "resource 1"},
				{"name": "resource 2"},
			})
			httpmock.RegisterResponder("GET", baseUrl+"/api/v1/resource", responder)
		})

		It("gets resources successfuly", func(ctx SpecContext) {
			resources, err := resourceClient.GetApiV1Resource(baseUrl, &pkg.Selector{})

			Expect(resources).To(Not(BeEmpty()))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	When("there is workflow with same name", func() {
	})
})

func TestResourceClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResourceClientSuite")
}
