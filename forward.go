package main

import (
	"fmt"

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
		fmt.Println("Couldn't create a private channel with", master.Username+".")
		fmt.Println("Channel :", c.Name)
		fmt.Println("Author :", m.Author.Username)
		fmt.Println("Message :", m.Content)
		fmt.Println(err.Error())
		return false
	}

	// Forward the message to Master!
	s.ChannelTyping(channel.ID)
	_, err = s.ChannelMessageSend(channel.ID, "<@"+m.Author.ID+"> : "+m.Content)
	if err != nil {
		fmt.Println("Couldn't forward a message to", master.Username+".")
		fmt.Println("Author :", m.Author.Username)
		fmt.Println("Message :", m.Content)
		fmt.Println(err.Error())
		return false
	}

	return true
}
