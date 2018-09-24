package main

import (
	"database/sql"
	"strings"

	"gitlab.com/NatoBoram/Go-Miiko/wheel"

	"github.com/bwmarrin/discordgo"
)

func sar(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, ms []string) {
	if len(ms) > 2 {
		sarCommand(s, g, c, m, strings.Join(ms[2:], " "))
	} else {

		// sar
		getSARsCommand(s, g, c, m)
	}
}

func sarCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	s.ChannelTyping(c.ID)

	// Recognize the role
	role, err := getRoleByString(s, g, roleString)
	if err != nil {
		printDiscordError("Couldn't get a role by its string.", g, c, m, nil, err)
		return
	}

	// Select it in the database
	_, err = selectSAR(g, role)
	if err == sql.ErrNoRows {
		_, err = s.ChannelMessageSend(c.ID, "Ce rôle n'est pas auto-assignable.")
		if err != nil {
			printDiscordError("Couldn't announce that a role isn't self-assignable.", g, c, m, nil, err)
		}
		return
	} else if err != nil {
		printDiscordError("Couldn't select a self-assignable role.", g, c, m, nil, err)
		return
	}

	// Get the member
	member, err := s.State.Member(g.ID, m.Author.ID)
	if err != nil {
		printDiscordError("Couldn't get a member.", g, c, m, nil, err)
		return
	}

	// Check if the user has the role
	if wheel.StringInSlice(role.ID, member.Roles) {

		// Remove the role
		err := s.GuildMemberRoleRemove(g.ID, member.User.ID, role.ID)
		if err != nil {
			printDiscordError("Couldn't un-assign a self-assignable role.", g, c, m, nil, err)

			// Feedback
			_, err = s.ChannelMessageSend(c.ID, "Désolée, mais je n'ai pas pu t'enlever ce rôle.")
			if err != nil {
				printDiscordError("Couldn't announce that I couldn't remove a self-assignable role.", g, c, m, nil, err)
			}
			return
		}

		// Removed!
		_, err = s.ChannelMessageSend(c.ID, "Je t'ai retiré le rôle **"+role.Name+"**.")
		if err != nil {
			printDiscordError("Couldn't announce that I removed a self-assignable role to someone.", g, c, m, nil, err)
		}
	} else {

		// Add the role
		err := s.GuildMemberRoleAdd(g.ID, member.User.ID, role.ID)
		if err != nil {
			printDiscordError("Couldn't assign a self-assignable role.", g, c, m, nil, err)

			// Feedback
			_, err = s.ChannelMessageSend(c.ID, "Désolée, mais je n'ai pas pu t'ajouter ce rôle.")
			if err != nil {
				printDiscordError("Couldn't announce that I couldn't add a self-assignable role.", g, c, m, nil, err)
			}
			return
		}

		// Assigned!
		_, err = s.ChannelMessageSend(c.ID, "Tu as maintenant le rôle **"+role.Name+"**.")
		if err != nil {
			printDiscordError("Couldn't announce that I added a self-assignable role to someone.", g, c, m, nil, err)
		}
	}
}
