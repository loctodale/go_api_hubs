package service

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/global"
	"strconv"
	"time"
)

type DiscordLogsService interface {
	SendMessage(channel string, message string) error
	SendTrackingLogs(channel string, message any, code int32, method string, path string) error
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

func (d *discordLogsService) SendTrackingLogs(channel string, message any, code int32, method string, path string) error {
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

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = discord.ChannelMessageSendEmbed(global.Config.DiscordBotLogs.Channel[channel], &discordgo.MessageEmbed{
		Title:       method + " " + path,
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0x00FF00,
		Description: "```json\n" + string(data) + "\n```",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Code",
				Value:  strconv.Itoa(int(code)),
				Inline: true,
			},
		},
	})

	return err
}
