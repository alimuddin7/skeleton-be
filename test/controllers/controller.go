package controllers

import (
	v1Controller "test/controllers/v1"
	"github.com/rs/zerolog"
)

type Controller interface {
	V1() v1Controller.V1Controller
}

type controller struct {
	V1Controller v1Controller.V1Controller
	Logs         zerolog.Logger
}

func InitializeController(v1 v1Controller.V1Controller, l zerolog.Logger) Controller {
	return &controller{
		V1Controller: v1,
		Logs:         l,
	}
}

func (c *controller) V1() v1Controller.V1Controller {
	return c.V1Controller
}
