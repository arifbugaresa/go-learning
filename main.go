package main

import (
	"go-learning/modules/user"
)

func main() {
	userRepo := user.NewUserRepository()
	userService := user.NewUserService(userRepo)

	GetMessageEndpointExampleForUnitTest(userService)
}

func GetMessageEndpointExampleForUnitTest(userService user.Service) (message string) {
	return userService.GetMessage()
}
