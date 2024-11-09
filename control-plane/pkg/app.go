package app

import (
	"github.com/kuzxnia/eris/control-plane/internal/infra"
	"github.com/kuzxnia/eris/control-plane/pkg/agent"
	"github.com/kuzxnia/eris/control-plane/pkg/config"
	"github.com/kuzxnia/eris/control-plane/pkg/interfaces/web"
	"github.com/kuzxnia/eris/control-plane/pkg/workflow"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	infra.Module,
  agent.Module,
	workflow.Module,
	web.Module,
)
