package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	HOST     = ""
	NAME     = ""
	USER     = ""
	PASSWORD = ""
	PORT     = ""
)

func LoadConfig() {
	HOST = viper.GetString("DB_HOST")
	NAME = viper.GetString("DB_NAME")
	USER = viper.GetString("DB_USER")
	PASSWORD = viper.GetString("DB_PASSWORD")
	PORT = viper.GetString("DB_PORT")
	log.Println(HOST, NAME, USER, PASSWORD, PORT)
}
