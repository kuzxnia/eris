package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kuzxnia/eris/control-plane/pkg/agent"
	"github.com/kuzxnia/eris/control-plane/pkg/web/dto"
)

type AgentController struct {
	service *agent.AgentService
}

func ProvideAgentController(service *agent.AgentService) *AgentController {
	return &AgentController{
		service: service,
	}
}

func (controller AgentController) RegisterRoutes(router fiber.Router) {
	agentRouter := router.Group("/agent")

	agentRouter.Post("/", controller.createAgent)
}

func (controller AgentController) createAgent(c *fiber.Ctx) error {
	agent := new(dto.AgentRequest)
	if err := c.BodyParser(agent); err != nil {
		return err
	}

	err := controller.service.CreateAgent(*agent)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.GeneralResponse{
		Message: "Created",
		Data:    agent,
	})
}
