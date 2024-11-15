package resourceselector_test

import (
	"github.com/kuzxnia/eris/control-plane/pkg/agent"
	resourceselector "github.com/kuzxnia/eris/control-plane/pkg/resource_selector"
	"github.com/kuzxnia/eris/control-plane/tests/unit/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Getting resources from agents", Ordered, Label("resource selector"), func() {
	var agentResourceSelector *resourceselector.AgentResourceSelector
	var agentRepositoryMock *mocks.MockAgentRepository
	var resourceClientMock *mocks.MockResourceClient
	BeforeEach(func(ctx SpecContext) {
		mockCtrl := gomock.NewController(GinkgoT())

		agentRepositoryMock = mocks.NewMockAgentRepository(mockCtrl)
		resourceClientMock = mocks.NewMockResourceClient(mockCtrl)
		agentResourceSelector = resourceselector.ProvideAgentResourceSelector(
			resourceClientMock, agentRepositoryMock,
		)
	})

	When("there are agents", func() {
		var agents []*agent.Agent
		var selector *resourceselector.Selector

		BeforeEach(func(ctx SpecContext) {
			selector = &resourceselector.Selector{}
			agents = []*agent.Agent{
				{URL: "http://agent-1.com"},
				{URL: "http://agent-2.com"},
			}
			agentRepositoryMock.
				EXPECT().GetAgents().Return(agents, nil)
			resourceClientMock.
				EXPECT().GetApiV1Resource("http://agent-1.com", selector).Return([]*resourceselector.Resource{{Name: "resource-from-agent1"}}, nil)
			resourceClientMock.
				EXPECT().GetApiV1Resource("http://agent-2.com", selector).Return([]*resourceselector.Resource{{Name: "resource-from-agent2"}}, nil)
		})

		It("gets resources from agents successfuly", func(ctx SpecContext) {
			resources, err := agentResourceSelector.Select(selector)

			Expect(resources).To(Not(BeEmpty()))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	When("there are no agents", func() {
		var selector *resourceselector.Selector
		BeforeEach(func(ctx SpecContext) {
			selector = &resourceselector.Selector{}
			agentRepositoryMock.
				EXPECT().GetAgents().Return([]*agent.Agent{}, nil)
		})

		It("gets empty list resources successfuly", func(ctx SpecContext) {
			resources, err := agentResourceSelector.Select(selector)

			Expect(resources).To(BeEmpty())
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
