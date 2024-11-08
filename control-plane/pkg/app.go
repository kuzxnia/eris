package app

import (
	"github.com/kuzxnia/eris/control-plane/pkg/interfaces/web"
	"github.com/kuzxnia/eris/control-plane/pkg/workflow"
	"go.uber.org/fx"
)

var Modules = fx.Options(
  workflow.Module,
	web.Module,
)
