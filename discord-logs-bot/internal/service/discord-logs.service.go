package service

import (
	"github.com/bwmarrin/discordgo"
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/global"
)

type DiscordLogsService interface {
	SendMessage(channel string, message string) error
}

type discordLogsService struct{}

func NewDiscordLogsService() DiscordLogsService {
	return &discordLogsService{}
}

func (d *discordLogsService) SendMessage(channel string, message string) error {
	discord, err := discordgo.New("Bot " + global.Config.DiscordBotLogs.Token)
	if err != nil {
		return err
	}
	discord.Identify.Intents = discordgo.IntentsGuildMessages
	err = discord.Open()
	if err != nil {
		return err
	}
	defer discord.Close()
	_, err = discord.ChannelMessageSend(global.Config.DiscordBotLogs.Channel[channel], message)
	return err
}
