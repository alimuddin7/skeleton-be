package usecases

import (
	"context"
	"github.com/rs/zerolog"
	hModels "test-new/helpers/models"
	"test-new/helpers"
)

type Usecase interface {
	HealthCheck(ctx context.Context) hModels.Response
	// [V1_USECASE_INTERFACE_MARKER]
}

type usecase struct {
	Logs zerolog.Logger
}

func InitializeV1Usecase(
	l zerolog.Logger,
) Usecase {
	return &usecase{
		Logs: l,
	}
}

func (u *usecase) HealthCheck(ctx context.Context) hModels.Response {
	var healthChecks []hModels.DataHealthCheck
	
	return helpers.GenerateResponseHealthCheck(healthChecks...)
}
