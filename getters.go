package main

import (
	"github.com/bwmarrin/discordgo"
)

func getWelcomeChannel(s *discordgo.Session, g *discordgo.Guild) (*discordgo.Channel, error) {

	// Select the presentation channel
	channelID, err := selectWelcomeChannel(g)
	if err != nil {
		return nil, err
	}

	// Turn it into a channel
	return s.Channel(channelID)
}

func getPresentationChannel(s *discordgo.Session, g *discordgo.Guild) (channel *discordgo.Channel, err error) {

	// Select the presentation channel
	channelID, err := selectPresentationChannel(g)
	if err != nil {
		return nil, err
	}

	// Turn it into a channel
	return s.Channel(channelID)
}
