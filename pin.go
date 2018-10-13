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
		s.UpdateStatus(0, "épingler "+m.Author.Username)

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
	s.UpdateStatus(0, "augmenter la difficulté de "+c.Name)

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
