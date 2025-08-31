package controllers

import "wallet-api/internal/services"

type TestController struct {
	service services.TestService
}

func (controller *TestController) Ping() string {
	return controller.service.Ping()
}
