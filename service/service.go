package service

import "legendary-kimchi/storage"

// Services contains all the repo structs
type Services struct {
	*ScooterService
}

// InitServices should be called in main.go
func InitServices(db storage.Storage) *Services {
	scooterRepo := NewScooterService(db)
	return &Services{ScooterService: scooterRepo}
}
