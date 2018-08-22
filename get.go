package main

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Get redirects the `get` coommand.
func Get(master *discordgo.User, db *sql.DB, s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, ms []string) {

	if len(ms) > 2 {
		switch ms[2] {
		case "welcome":
			// Get Welcome Channel
			if len(ms) > 3 {
				if ms[3] == "channel" {
					getWelcomeChannelCommand(s, g, c)
				}
			}
			break
		case "presentation":
			// Get Presentation Channel
			if len(ms) > 3 {
				if ms[3] == "channel" {
					getPresentationChannelCommand(s, g, c)
				}
			}
			break
		case "points":
			// Get Points
			GetPoints(s, g, c, m)
			break
		case "lover":
			// Get Lover
			// GetLoverCmd(db, s, g, c, m.Author)
			break
		}
	}
}

// GetWelcomeChannelCommand send the welcome channel to an user.
func getWelcomeChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {

	// Get the welcome channel
	channel, err := getWelcomeChannel(s, g)
	if err != nil {
		s.ChannelMessageSend(c.ID, "Il n'y a pas de salon de bienvenue.")
		return
	}

	// Send the welcome channel
	s.ChannelMessageSend(c.ID, "Le salon de bienvenue est <#"+channel.ID+">.")
	if err != nil {
		fmt.Println("Couldn't send the welcome channel.")
		fmt.Println("Guild : " + g.Name)
		fmt.Println("Channel : " + c.Name)
		fmt.Println(err.Error())
		return
	}
}

func getPresentationChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {

	// Get the presentation channel
	channel, err := getPresentationChannel(s, g)
	if err != nil {
		s.ChannelMessageSend(c.ID, "Il n'y a pas de salon de présentation.")
		return
	}

	s.ChannelMessageSend(c.ID, "Le salon de présentation est <#"+channel.ID+">.")
	if err != nil {
		fmt.Println("Couldn't send the presentation channel.")
		fmt.Println("Guild : " + g.Name)
		fmt.Println("Channel : " + c.Name)
		fmt.Println(err.Error())
		return
	}
}
