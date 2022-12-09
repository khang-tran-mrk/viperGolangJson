package main

import (
	"TestViperJson/configs"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("dev")  // Register config file name (no extension)
	viper.SetConfigType("json") // Look for specific type
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	config := configs.Config{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Println(err)
	} else {
		log.Println(config)
	}

}
