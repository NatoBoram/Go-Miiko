package main

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
)

func setPresentationChannel(g *discordgo.Guild, c *discordgo.Channel) (sql.Result, error) {

	// Check if there's a presentation channel
	_, err := selectPresentationChannel(g)
	if err == sql.ErrNoRows {

		// Insert if there's none
		return insertPresentationChannel(g, c)
	} else if err != nil {
		return nil, err
	}

	// Update if there's one
	return updatePresentationChannel(g, c)
}

func setWelcomeChannel(g *discordgo.Guild, c *discordgo.Channel) (sql.Result, error) {

	// Check if there's a presentation channel
	_, err := selectWelcomeChannel(g)
	if err == sql.ErrNoRows {

		// Insert if there's none
		return insertWelcomeChannel(g, c)
	} else if err != nil {
		return nil, err
	}

	// Update if there's one
	return updateWelcomeChannel(g, c)
}
