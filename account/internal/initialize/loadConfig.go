package initialize

import (
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/account/global"
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

	if err = viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}
	fmt.Println("Config Server port::", global.Config.AccountService.Ports.Local)
}
