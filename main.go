package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-learning/utils/configs"
)

func main() {
	configs.ReadConfigurations()

	fmt.Println(viper.GetString("name"))
	fmt.Println(viper.GetString("app.mode"))
}
