package main

import (
	"fmt"

	"github.com/spf13/viper"
)

var ()

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config file not found in working dir"))
	}

	token := viper.GetString("discord_token")

	fmt.Println(token)
}
