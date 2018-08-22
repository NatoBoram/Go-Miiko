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
