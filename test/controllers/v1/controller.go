package controllers

import (
	"reflect"
	"strings"
	v1Usecases "test/usecases/v1"
	"github.com/gofiber/fiber/v3"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"net/http"
)

type V1Controller interface {
	HealthCheck(c fiber.Ctx) error
	// [V1_CONTROLLER_INTERFACE_MARKER]
}

type v1Controller struct {
	Usecase   v1Usecases.Usecase
	Logs      zerolog.Logger
	Validator *validator.Validate
}

func InitializeV1Controller(usecases v1Usecases.Usecase, l zerolog.Logger) V1Controller {
	return &v1Controller{
		Usecase:   usecases,
		Logs:      l,
		Validator: CustomValidator(),
	}
}

// CustomValidator creates a new validator instance with custom tag name function
func CustomValidator() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	// register custom validation here
	// example: v.RegisterValidation("amount", helpers.IsValidFloat)
	
	return v
}

func (ctrl *v1Controller) HealthCheck(c fiber.Ctx) error {
	res := ctrl.Usecase.HealthCheck(c.Context())
	return c.Status(http.StatusOK).JSON(res)
}
