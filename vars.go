package main

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
)

var (
	db     *sql.DB
	me     *discordgo.User
	master *discordgo.User
)

var popcornStrings = [...]string{
	"popcorn",
	"pop-corn",
	"pop corn",
	"maïs soufflé",
	"maïs éclaté",
	"🍿",
}

var censoredUsernames = [...]string{
	"discord.gg",
	"twitch.tv",
	"twitter.com",
}
