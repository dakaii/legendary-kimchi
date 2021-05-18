package controller

import (
	"legendary-kimchi/controller/scooter"
	"legendary-kimchi/service"
)

// Controllers contains all the controllers
type Controllers struct {
	scooterController *scooter.Controller
}

// InitControllers returns a new Controllers
func InitControllers(services *service.Services) *Controllers {
	return &Controllers{
		scooterController: scooter.InitController(services.ScooterService),
	}
}
