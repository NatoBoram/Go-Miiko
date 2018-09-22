package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func addHandlers(s *discordgo.Session) {

	s.AddHandler(messageHandler)
	s.AddHandler(reactHandler)
	s.AddHandler(leaveHandler)
	s.AddHandler(joinHandler)

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Myself? Super User?
	if m.Author.ID == me.ID || m.Author.Discriminator == "0000" {
		return
	}

	// Flow Control
	done := false

	// Get channel structure
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get a channel structure.")
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	// Forward to Master.
	done = forward(s, channel, m.Message)
	if done {
		return
	}

	// Get guild structure
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Couldn't get a guild structure.")
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	// Get guild member
	member, err := s.GuildMember(channel.GuildID, m.Author.ID)
	if err != nil {
		fmt.Println("Couldn't get a member structure.")
		fmt.Println("Guild : " + guild.Name)
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	// Update welcome channel
	if m.Type == discordgo.MessageTypeGuildMemberJoin {
		setWelcomeChannel(guild, channel)
		return
	}

	// Guard
	done = placeInAGuard(s, guild, channel, member, m.Message)
	if done {
		return
	}

	// Nani?!
	done = Nani(s, m.Message)
	if done {
		return
	}

	// Popcorn!
	done = Popcorn(s, channel, m.Message)
	if done {
		return
	}

	// Mentionned someone?
	if len(m.Mentions) == 1 {

		// Mentionned me?
		if m.Mentions[0].ID == me.ID {

			// Split
			command := strings.Split(m.Content, " ")

			// Redirect commands
			if len(command) > 1 {
				switch command[1] {
				case "prune":
					prune(s, guild, channel, m.Message)
					return
				case "get":
					get(s, guild, channel, m.Message, command)
					return
				case "set":
					set(s, guild, channel, m.Message, command)
					return
				}
			}
		}
	}

	// Love!
	// done = love(s, guild, channel, m.Message)
	if done {
		return
	}
}

func reactHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Get the message structure
	message, err := s.ChannelMessage(m.ChannelID, m.MessageID)
	if err != nil {
		fmt.Println("Couldn't get the message structure of a MessageReactionAdd!")
		fmt.Println("ChannelID : " + m.ChannelID)
		fmt.Println(err.Error())
		return
	}

	// Get channel structure
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get the channel structure of a MessageReactionAdd!")
		fmt.Println("ChannelID : " + m.ChannelID)
		fmt.Println("Author : " + message.Author.Username)
		fmt.Println("Message : " + message.Content)
		fmt.Println(err.Error())
		return
	}

	// Get the guild structure
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Couldn't get the guild structure of a MessageReactionAdd!")
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + message.Author.Username)
		fmt.Println("Message : " + message.Content)
		fmt.Println(err.Error())
		return
	}

	// Pin popular message
	pin(s, guild, channel, message)
}

func leaveHandler(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

	// Get guild
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Couldn't get the guild of " + m.User.Username + "!")
		fmt.Println(err.Error())
		return
	}

	// Invite people who leave
	waitComeBack(s, guild, m.Member)
}

func joinHandler(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	// Get guild
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Couldn't get the guild", m.User.Username, "just joined.")
		fmt.Println(err.Error())
		return
	}

	// Ask for guard
	askForGuard(s, guild, m.Member)
}
