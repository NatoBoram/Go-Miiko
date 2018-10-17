package main

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/NatoBoram/Go-Miiko/wheel"
)

func pin(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message) {

	// DM?
	if c.Type == discordgo.ChannelTypeDM {
		return
	}

	// Get the reactions
	var singleReactionCount int
	for _, reaction := range m.Reactions {
		singleReactionCount = wheel.MaxInt(singleReactionCount, reaction.Count)
	}

	// Minimum reactions
	minReactions, err := getMinimumReactions(g, c)
	if err != nil {
		printDiscordError("Couldn't get the minimum reactions for a channel.", g, c, m, nil, err)
	}

	// Check the reactions
	if singleReactionCount >= minReactions {

		// Pin it!
		err = s.ChannelMessagePin(c.ID, m.ID)
		if err != nil {
			printDiscordError("Couldn't pin a popular message!", g, c, m, nil, err)

			// Check the amount of pins in that channel
			messages, err := s.ChannelMessagesPinned(c.ID)
			if err != nil {
				printDiscordError("Couldn't obtain the amount of pins in a channel.", g, c, m, nil, err)
				return
			}

			// Upgrade the minimum
			if len(messages) >= 50 {
				err = addMinimumReactions(c)
				if err != nil {
					printDiscordError("Couldn't add to the minimum reactions of a channel", g, c, m, nil, err)
					return
				}

				purgePin(s, g, c, m, messages)
			}

			return
		}

		// Check if already in the database
		_, err := selectPin(m)
		if err == sql.ErrNoRows {

			// Status
			err = setStatus(s, "épingler "+m.Author.Username)
			if err != nil {
				printDiscordError("Couldn't set the status to pinning someone.", g, c, m, nil, err)
			}

			// Throw it in the hall of fame
			savePin(s, g, m)
			if err == sql.ErrNoRows {
				// There's no hall of fame in this server
			} else if err != nil {
				printDiscordError("Couldn't throw a message in the hall of fame.", g, c, m, nil, err)
			}

			// Not previously pinned, time to insert it!
			_, err = insertPin(g, m)
			if err != nil {
				printDiscordError("Couldn't insert a pin.", g, c, m, nil, err)
			}
		} else if err != nil {
			printDiscordError("Couldn't select a pin.", g, c, m, nil, err)
		}
	}
}

func purgePin(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, messages []*discordgo.Message) {
	err := setStatus(s, "augmenter la difficulté de "+c.Name)
	if err != nil {
		printDiscordError("Couldn't set the status to upping difficulty for a channel.", g, c, m, nil, err)
	}

	// Check minimum
	channelMin, err := selectMinimumReactions(c)
	if err != nil {
		printDiscordError("Couldn't add to the minimum reactions of a channel", g, c, m, nil, err)
	}

	// For each messages
	for _, message := range messages {

		// Check pins
		var singleReactionCount int
		for _, reaction := range message.Reactions {
			singleReactionCount = wheel.MaxInt(singleReactionCount, reaction.Count)
		}

		// Unpin
		if channelMin > singleReactionCount {
			err := s.ChannelMessageUnpin(c.ID, message.ID)
			if err != nil {
				printDiscordError("Couldn't unpin a previously popular message", g, c, message, nil, err)
				continue
			}
		}

		// Delete pin
		_, err := deletePin(message)
		if err != nil {
			printDiscordError("Couldn't remove a pin from the database.", g, c, message, nil, err)
		}
	}
}

func savePin(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Message) (message *discordgo.Message, err error) {

	// Get the channel
	halloffame, err := getFameChannel(s, g)
	if err != nil {
		return
	}

	// Don't ping the author
	message, err = s.ChannelMessageSend(halloffame.ID, "**"+m.Author.Username+" :** "+m.Content)
	if err != nil {
		return
	}

	// Change to a mention
	message, err = s.ChannelMessageEdit(halloffame.ID, message.ID, m.Author.Mention()+" : "+m.Content)
	if err != nil {
		return
	}

	return
}
