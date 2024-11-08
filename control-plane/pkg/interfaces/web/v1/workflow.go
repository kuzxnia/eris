package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kuzxnia/eris/control-plane/pkg/interfaces/web/dto"
	"github.com/kuzxnia/eris/control-plane/pkg/workflow"
)

type WorkflowController struct {
	service *workflow.WorkflowService
}

func ProvideWorkflowController(service *workflow.WorkflowService) *WorkflowController {
	return &WorkflowController{
		service: service,
	}
}

func (controller WorkflowController) RegisterRoutes(router fiber.Router) {
	workflowRouter := router.Group("/workflow")

	workflowRouter.Post("/", controller.createWorkflow)
}

func (controller WorkflowController) createWorkflow(c *fiber.Ctx) error {
	workflow := new(dto.WorkflowRequest)
	if err := c.BodyParser(workflow); err != nil {
		return err
	}

  err := controller.service.CreateWorkflow(*workflow)

  if err != nil {
    return err
  }

	return c.Status(fiber.StatusOK).JSON(dto.GeneralResponse{
		Message: "Created",
		Data:    workflow,
	})
}
