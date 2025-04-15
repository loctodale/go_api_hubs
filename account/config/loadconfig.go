package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//fmt.Println("Server port::", viper.GetInt("server.port"))
	//fmt.Println("Jwt Key::", viper.GetString("security.jwt.certs"))
}
