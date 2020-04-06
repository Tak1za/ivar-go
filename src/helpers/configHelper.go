package helpers

import (
	"github.com/spf13/viper"
	"ivar-go/src/models"
	"log"
)

var Config models.Configuration

func ReadConfiguration() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("Config file not found!")
		} else {
			log.Fatalln("Config file found, not able to parse it!")
		}
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %s", err)
	}
}
