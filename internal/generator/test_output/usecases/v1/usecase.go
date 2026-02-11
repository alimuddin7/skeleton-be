package usecases

import (
	"context"
	"github.com/rs/zerolog"
	hModels "test_output/helpers/models"
	"test_output/helpers"
	"test_output/databases"
)

type Usecase interface {
	HealthCheck(ctx context.Context) hModels.Response
	// [V1_USECASE_INTERFACE_MARKER]
}

type usecase struct {
	DB databases.Database
	Logs zerolog.Logger
}

func InitializeV1Usecase(
	db databases.Database,
	l zerolog.Logger,
) Usecase {
	return &usecase{
		DB: db,
		Logs: l,
	}
}

func (u *usecase) HealthCheck(ctx context.Context) hModels.Response {
	var healthChecks []hModels.DataHealthCheck
	// Database health checks
	healthChecks = append(healthChecks, u.DB.GetMysql().HealthCheck(ctx))
	
	return helpers.GenerateResponseHealthCheck(healthChecks...)
}
