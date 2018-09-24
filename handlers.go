package main

import (
	"database/sql"
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
		printDiscordError("Couldn't get a member structure.", guild, channel, m.Message, nil, err)
		return
	}

	// First time setup for welcome channels
	if m.Type == discordgo.MessageTypeGuildMemberJoin {

		// Check if there's already one
		_, err := selectWelcomeChannel(guild)
		if err == sql.ErrNoRows {

			// Set the welcome channel
			_, err := setWelcomeChannel(guild, channel)
			if err != nil {
				printDiscordError("Couldn't select the welcome channel.", guild, channel, m.Message, nil, err)
			}

			// Ask for guard, because the joinHandler probably couldn't handle it
			askForGuard(s, guild, m.Message.Author)
		} else if err != nil {
			printDiscordError("Couldn't select the welcome channel.", guild, channel, m.Message, nil, err)
		}
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

					// Permissions
					has, err := canAdministrate(s, guild, member)
					if err != nil {
						printDiscordError("Couldn't check if a member can administrate.", guild, channel, m.Message, nil, err)
						return
					}

					// Check
					if !has {
						s.ChannelMessageSend(channel.ID, "Tu n'as pas la permission de faire ça.")
						return
					}

					prune(s, guild, channel, m.Message)
					return
				case "get":

					// Permissions
					has, err := isGuardianOrOver(s, guild, member)
					if err != nil {
						printDiscordError("Couldn't check if a member has a guard.", guild, channel, m.Message, nil, err)
						return
					}

					// Check
					if !has {
						s.ChannelMessageSend(channel.ID, "Tu n'as pas la permission de faire ça.")
						return
					}

					get(s, guild, channel, m.Message, command)
					return
				case "set":

					// Permissions
					has, err := canAdministrate(s, guild, member)
					if err != nil {
						printDiscordError("Couldn't check if a member can administrate.", guild, channel, m.Message, nil, err)
						return
					}

					// Check
					if !has {
						s.ChannelMessageSend(channel.ID, "Tu n'as pas la permission de faire ça.")
						return
					}

					set(s, guild, channel, m.Message, command)
					return
				case "info":

					// Permissions
					has, err := canModerate(s, guild, member)
					if err != nil {
						printDiscordError("Couldn't check if a member can moderate.", guild, channel, m.Message, nil, err)
						return
					}

					// Check
					if !has {
						s.ChannelMessageSend(channel.ID, "Tu n'as pas la permission de faire ça.")
						return
					}

					info(s, guild, channel, m.Message, command)
					return
				case "sar":

					// Permissions
					has, err := hasGuard(s, guild, member)
					if err != nil {
						printDiscordError("Couldn't check if a member has a guard.", guild, channel, m.Message, nil, err)
						return
					}

					// Check
					if !has {
						s.ChannelMessageSend(channel.ID, "Tu n'as pas la permission de faire ça.")
						return
					}

					sar(s, guild, channel, m.Message, command)
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

	// Don't announce it if its username contains "discord.gg".
	if strings.Contains(strings.ToLower(m.User.Username), "discord.gg") {
		return
	}

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

	// Ban those whose username contains "discord.gg".
	if strings.Contains(strings.ToLower(m.User.Username), "discord.gg") {
		err := s.GuildBanCreateWithReason(m.GuildID, m.User.ID, "Lien d'invitation dans le username.", 7)
		if err != nil {
			printDiscordError("Couldn't ban a bot spammer.", nil, nil, nil, m.User, err)
		}
		return
	}

	// Get guild
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		printDiscordError("Couldn't get the guild "+m.User.Username+" just joined.", guild, nil, nil, m.User, err)
		return
	}

	// Ask for guard
	askForGuard(s, guild, m.Member.User)
}
