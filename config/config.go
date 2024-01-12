package config

import (
	"github.com/spf13/viper"
)

var (
	HOST              = ""
	NAME              = ""
	USER              = ""
	PASSWORD          = ""
	PORT              = ""
	XML_PATH          = ""
	ZIP_PATH          = ""
	XML_NAME          = ""
	ZIP_NAME          = ""
	JSON_FILE_PATH    = ""
	ES_HOST           = ""
	ES_INDEX_NAME     = ""
	XML_TARGET_FOLDER = ""
)

func LoadConfig() {
	HOST = viper.GetString("DB_HOST")
	NAME = viper.GetString("DB_NAME")
	USER = viper.GetString("DB_USER")
	PASSWORD = viper.GetString("DB_PASSWORD")
	PORT = viper.GetString("DB_PORT")

	XML_PATH = viper.GetString("XML_PATH")
	ZIP_PATH = viper.GetString("ZIP_PATH")
	XML_NAME = viper.GetString("XML_NAME")
	ZIP_NAME = viper.GetString("ZIP_NAME")
	JSON_FILE_PATH = viper.GetString("JSON_FILE_PATH")
	XML_TARGET_FOLDER = viper.GetString("XML_TARGET_FOLDER")

	ES_HOST = viper.GetString("ES_HOST")
	ES_INDEX_NAME = viper.GetString("ES_INDEX_NAME")

}
