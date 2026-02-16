package helpers

import (
	"test/models"
	"net/http"
)

func GenerateResponseHealthCheck(data ...models.DataHealthCheck) models.Response {
	return models.Response{
		Meta: models.MetaData{
			Code:    "2000000",
			Title:   "Available",
			Message: "Service is available",
		},
		Data: data,
	}
}
