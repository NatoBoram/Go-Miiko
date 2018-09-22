package main

import (
	"github.com/bwmarrin/discordgo"
)

// Database hosts the bot's database configuration.
type Database struct {
	User     string
	Password string
	Address  string
	Port     int
	Database string
}

// Discord hosts the bot's Discord configuration.
type Discord struct {
	Token    string
	MasterID string
}

// Languages is the structure that holds all the bot's supported languages.
type Languages struct {
	French  string
	English string
}

type discordgoStackTrace struct {
	info    string
	guild   *discordgo.Guild
	channel *discordgo.Channel
	message *discordgo.Message
	user    *discordgo.User
	err     error
}
