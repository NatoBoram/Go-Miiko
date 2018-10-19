package main

import (
	"database/sql"
	"time"

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
		_, _, err := selectPin(m)
		if err == sql.ErrNoRows {

			// Status
			err = setStatus(s, "épingler "+m.Author.Username)
			if err != nil {
				printDiscordError("Couldn't set the status to pinning someone.", g, c, m, nil, err)
			}

			// Throw it in the hall of fame
			go savePin(s, g, m)

			// Not previously pinned, time to insert it!
			_, err = insertPin(g, c, m)
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

func savePin(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Message) (saved bool) {

	// Only text messages are transferred. Delete empty messages.
	if m.Type != discordgo.MessageTypeDefault || m.Content == "" {
		return
	}

	// Get the hall of fame
	halloffame, err := getFameChannel(s, g)
	if err != nil || m.ChannelID == halloffame.ID {
		printDiscordError("Couldn't get the hall of fame.", g, nil, m, nil, err)
		return
	}

	// Just in case `GuildMember` fails
	var (
		colour = colourNPC
		name   = m.Author.Username
	)

	// Get Member
	member, err := s.GuildMember(g.ID, m.Author.ID)
	if err != nil {
		// Probably just a member that's dead. Nothing to output.
		// printDiscordError("Couldn't get a pinned member.", g, nil, m, nil, err)
	} else {

	// Get colour
		colour, _ = getColour(s, g, member)

	// Get name
		if member == nil {
			name = m.Author.Username
		} else if member.Nick == "" {
			name = member.User.Username
	} else {
			name = member.Nick
	}
	}

	// Create Embed
	embed := &discordgo.MessageEmbed{
		Color: colour,
		Author: &discordgo.MessageEmbedAuthor{
			URL:     "https://canary.discordapp.com/channels/" + g.ID + "/" + m.ChannelID + "/" + m.ID + "/",
			Name:    name,
			IconURL: m.Author.AvatarURL(""),
		},
		URL:         "https://canary.discordapp.com/channels/" + g.ID + "/" + m.ChannelID + "/" + m.ID + "/",
		Title:       "Message",
		Description: m.Content,
	}

	var (
		smallest *discordgo.MessageAttachment
		largest  *discordgo.MessageAttachment
	)

	// For all attachments
	for _, attachment := range m.Attachments {

		// Fields
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "Attachement",
			Value:  "[" + attachment.Filename + "](" + attachment.URL + ")",
			Inline: true,
		})

		// Check if image
		if attachment.Height == 0 || attachment.Width == 0 {
			continue
		}

		// Play with sizes
		if smallest == nil && largest == nil {
			if attachment.Width > attachment.Height || attachment.Width > 300 {
				largest = attachment
			} else {
				smallest = attachment
			}
		} else if smallest == nil {
			if attachment.Height*attachment.Width > largest.Height*largest.Width {
				smallest, largest = largest, attachment
			} else {
				smallest = attachment
			}
		} else if largest == nil {
			if attachment.Height*attachment.Width < smallest.Height*smallest.Width {
				smallest, largest = attachment, smallest
			} else {
				largest = attachment
			}
		} else {
			if attachment.Height*attachment.Width < smallest.Height*smallest.Width {
				smallest = attachment
			} else if attachment.Height*attachment.Width > largest.Height*largest.Width {
				largest = attachment
			}
		}
	}

	// Thumbnail. 80x80.
	if smallest != nil {
		embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL:    smallest.URL,
			Width:  smallest.Width,
			Height: smallest.Height,
		}
	}

	// Image. 400x300.
	if largest != nil {
		embed.Image = &discordgo.MessageEmbedImage{
			URL:    largest.URL,
			Width:  largest.Width,
			Height: largest.Height,
		}
	}

	// Channel
	embed.Fields = append(embed.Fields,
		&discordgo.MessageEmbedField{
			Name:   "Salon",
			Value:  "<#" + m.ChannelID + ">",
			Inline: true,
		},
	)

	// Author
	embed.Fields = append(embed.Fields,
		&discordgo.MessageEmbedField{
			Name:   "Auteur",
			Value:  "<@" + m.Author.ID + ">",
			Inline: true,
		},
	)

	// Emoji Field
	if len(m.Reactions) > 0 {
		emojiField := &discordgo.MessageEmbedField{
			Name: "Réactions",
		}

		// For each reactions
		for _, reaction := range m.Reactions {

			// Check if RequireColons
			if reaction.Emoji.ID != "" && reaction.Emoji.Name != "" {
				emojiField.Value += "<:" + reaction.Emoji.Name + ":" + reaction.Emoji.ID + ">"
			} else if reaction.Emoji.Name != "" {
				emojiField.Value += reaction.Emoji.Name
			} else {
				emojiField.Value += "<:" + reaction.Emoji.ID + ">"
			}
		}
		embed.Fields = append(embed.Fields, emojiField)
	}

	// Footer
	embed.Footer = &discordgo.MessageEmbedFooter{
		IconURL: discordgo.EndpointGuildIcon(g.ID, g.Icon),
		Text:    wheel.ToFrenchDate(time.Now()),
	}

	// Send embed
	_, err = s.ChannelMessageSendEmbed(halloffame.ID, embed)
	if err != nil {
		printDiscordError("Couldn't send an embed.", g, nil, m, nil, err)
		return
	}

	// Save it in the database
	_, err = insertMessagesFamed(g, m)
	if err != nil {
		printDiscordError("Couldn't insert a famed message in the hall of fame!", g, nil, m, nil, err)
		return
	}

	return true
}
