package main

import (
	"fmt"
	"go-learning/modules/user"
)

func main() {
	userRepo := user.NewUserRepository()
	userService := user.NewUserService(userRepo)

	message := userService.GetMessage()

	fmt.Println(message)
}
