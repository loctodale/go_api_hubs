package initialize

import (
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/auth/global"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	viper.AddConfigPath("./config")
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err = viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}
	fmt.Println("Config Server port::", global.Config.AuthService.Ports.Local)
}
