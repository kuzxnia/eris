package web

import (
	"context"

	"github.com/kuzxnia/eris/agent/pkg/interfaces/web/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkflowController struct {
	proto.UnimplementedWorkflowControllerServer
}

func ProvideWorkflowController() *WorkflowController {

	return &WorkflowController{}
}

// nie ma potrzeby robienia dto bo proto to wałaściwie dto??
func (WorkflowController) RunWorkflow(context.Context, *proto.RunWorkflowRequest) (*proto.RunWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunWorkflow not implemented")
}

func (WorkflowController) StopWorkflow(context.Context, *proto.StopWorkflowRequest) (*proto.StopWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopWorkflow not implemented")
}

func (controller *WorkflowController) Register(s grpc.ServiceRegistrar) error {
	s.RegisterService(&proto.WorkflowController_ServiceDesc, controller)
	return nil
}
