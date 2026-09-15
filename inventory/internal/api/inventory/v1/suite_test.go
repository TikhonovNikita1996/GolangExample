package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/service/mocks"
)

type ApiSuite struct {
	suite.Suite

	ctx context.Context

	service *mocks.InventoryService

	api *api
}

func (s *ApiSuite) SetupTest() {
	s.ctx = context.Background()

	s.service = mocks.NewInventoryService(s.T())

	s.api = NewAPI(
		s.service,
	)
}

func (s *ApiSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ApiSuite))
}
