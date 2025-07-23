package initialize

import (
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/global"
	"github.com/spf13/viper"
)

func LoadDiscordConfig() {
	env := "docker"
	viper.AddConfigPath("/etc/discord-logs-bot/")
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err = viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}
	fmt.Println(global.Config.DiscordBotLogs.Token)
	fmt.Println("Config Server port::", 8080)
}
