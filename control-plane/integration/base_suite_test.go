package integration

import (
	"context"
	"testing"

	"github.com/gofiber/fiber/v2"
	main "github.com/kuzxnia/eris/control-plane/pkg"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
)

var (
	FxApp    *fx.App
	FiberApp *fiber.App
)

type TestSuite struct {
	suite.Suite
	FxApp    *fx.App
	FiberApp *fiber.App
}

func (s *TestSuite) SetupSuite() {
	if FxApp == nil {
		s.FxApp = fx.New(
			main.Modules,
			// fx.Replace(mocks(t)),
			// fx.Invoke(r),
			fx.Invoke(func(fiberApp *fiber.App) {
				s.FiberApp = fiberApp
			}),
		)
		FxApp = s.FxApp
		err := s.FxApp.Start(context.Background())
		s.Require().NoError(err)
	}
}

func (s *TestSuite) TearDownSuite() {
	if FxApp != nil {
		err := s.FxApp.Stop(context.Background())
		s.Require().NoError(err)
	}
}

func TestSuites(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
