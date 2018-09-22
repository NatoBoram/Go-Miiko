package main

import (
	"github.com/bwmarrin/discordgo"
)

func forward(s *discordgo.Session, c *discordgo.Channel, m *discordgo.Message) bool {

	// DM, Master, Me
	if c.Type != discordgo.ChannelTypeDM || m.Author.ID == master.ID || m.Author.ID == me.ID {
		return false
	}

	// Create channel with Master
	channel, err := s.UserChannelCreate(master.ID)
	if err != nil {
		printDiscordError("Couldn't create a private channel with"+master.Username+".", nil, c, m, master, err)
		return false
	}

	// Forward the message to Master!
	s.ChannelTyping(channel.ID)
	_, err = s.ChannelMessageSend(channel.ID, "<@"+m.Author.ID+"> : "+m.Content)
	if err != nil {
		printDiscordError("Couldn't forward a message to"+master.Username+".", nil, channel, m, master, err)
		return false
	}

	return true
}
