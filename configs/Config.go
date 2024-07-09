package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitiateConfiguration() {
	viper.SetConfigName("Config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	fmt.Println("Successfully read config file")
}
