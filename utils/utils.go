package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func ImportENV() {

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Print("Error loading .env file")
	}

}
