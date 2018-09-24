package main

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func info(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, ms []string) {

	// info
	if len(ms) > 3 {
		switch ms[2] {
		case "channel":
			infoChannelCommand(s, g, c, m, ms[3])
		case "member":
			infoMemberCommand(s, g, c, m, ms[3])

		// info ?
		default:
		}
	} else {
		// info

	}
}

func infoChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, id string) {
	s.ChannelTyping(c.ID)

	channel, err := s.State.Channel(id)
	if err != nil {
		printDiscordError("Couldn't get the specified channel.", g, c, m, nil, err)
		return
	}

	// Create Embed
	embed := &discordgo.MessageEmbed{
		Color:  colourBot,
		Fields: []*discordgo.MessageEmbedField{},
	}

	// embed.Fields = addEmbedField(embed.Fields, "Bitrate", strconv.Itoa(channel.Bitrate), true)
	// embed.Fields = addEmbedField(embed.Fields, "GuildID", channel.GuildID, true)
	embed.Fields = addEmbedField(embed.Fields, "ID", channel.ID, true)
	// embed.Fields = addEmbedField(embed.Fields, "LastMessageID", channel.LastMessageID, true)
	// embed.Fields = addEmbedField(embed.Fields, "Messages", strconv.Itoa(len(channel.Messages)), true)
	embed.Fields = addEmbedField(embed.Fields, "Name", channel.Name, true)
	embed.Fields = addEmbedField(embed.Fields, "NSFW", strconv.FormatBool(channel.NSFW), true)
	// embed.Fields = addEmbedField(embed.Fields, "ParentID", channel.ParentID, true)
	// embed.Fields = addEmbedField(embed.Fields, "PermissionOverwrites", strconv.Itoa(len(channel.PermissionOverwrites)), true)
	// embed.Fields = addEmbedField(embed.Fields, "Position", strconv.Itoa(channel.Position), true)
	// embed.Fields = addEmbedField(embed.Fields, "Recipients", strconv.Itoa(len(channel.Recipients)), true)
	embed.Fields = addEmbedField(embed.Fields, "Topic", channel.Topic, true)

	// Send embed
	_, err = s.ChannelMessageSendEmbed(c.ID, embed)
	if err != nil {
		printDiscordError("Couldn't send an embed.", g, c, m, nil, err)
	}
}

func addEmbedField(fields []*discordgo.MessageEmbedField, name string, value string, inline bool) []*discordgo.MessageEmbedField {
	if value != "" {
		return append(fields, &discordgo.MessageEmbedField{
			Name:   name,
			Value:  value,
			Inline: inline,
		})
	}
	return fields
}

func infoMemberCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, id string) {
	s.ChannelTyping(c.ID)

	// Get the member
	member, err := s.GuildMember(g.ID, id)
	if err != nil {
		printDiscordError("Couldn't get a guild's member.", g, c, m, nil, err)
		return
	}

	// Get highest role
	colour, err := getColour(s, g, member)
	if err != nil {
		printDiscordError("Couldn't get a member's colour", g, c, m, nil, err)
	}

	// Create Embed
	embed := &discordgo.MessageEmbed{
		Color: colour,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: member.User.AvatarURL(""),
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    member.User.Username,
			IconURL: member.User.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{},
	}

	// Member
	embed.Fields = addEmbedField(embed.Fields, "Joined at", member.JoinedAt, true)
	embed.Fields = addEmbedField(embed.Fields, "Nickname", member.Nick, true)
	// embed.Fields = addEmbedField(embed.Fields, "Roles", strings.Join(member.Roles, ", "), true)

	// User
	embed.Fields = addEmbedField(embed.Fields, "Bot", strconv.FormatBool(member.User.Bot), true)
	embed.Fields = addEmbedField(embed.Fields, "Email", member.User.Email, true)
	embed.Fields = addEmbedField(embed.Fields, "ID", member.User.ID, true)
	embed.Fields = addEmbedField(embed.Fields, "Mention", member.User.Mention(), true)
	embed.Fields = addEmbedField(embed.Fields, "Multiple Factor Authentication", strconv.FormatBool(member.User.MFAEnabled), true)
	embed.Fields = addEmbedField(embed.Fields, "String", member.User.String(), true)
	embed.Fields = addEmbedField(embed.Fields, "Token", member.User.Token, true)
	embed.Fields = addEmbedField(embed.Fields, "Verified", strconv.FormatBool(member.User.Verified), true)

	// Send embed
	_, err = s.ChannelMessageSendEmbed(c.ID, embed)
	if err != nil {
		printDiscordError("Couldn't send an embed.", g, c, m, nil, err)
	}
}
