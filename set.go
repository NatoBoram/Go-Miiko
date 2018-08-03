package main

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Set redirects the `set` coommand.
func Set(db *sql.DB, s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, ms []string) {

	if len(ms) > 2 {
		switch ms[2] {
		case "welcome":
			// Set Welcome Channel
			if len(ms) > 3 {
				if ms[3] == "channel" {
					if m.Author.ID == g.OwnerID {
						setWelcomeChannelCommand(g, c)
					}
				}
			}
			break
		case "presentation":
			// Set Presentation Channel
			if len(ms) > 3 {
				if ms[3] == "channel" {
					if m.Author.ID == g.OwnerID {
						setPresentationChannelCommand(g, c)
					}
				}
			}
			break
		}
	}
}

// setWelcomeChannelCommand sets the welcome channel and sends feedback to the user.
func setWelcomeChannelCommand(g *discordgo.Guild, c *discordgo.Channel) {

	// Set the welcome channel
	_, err := setWelcomeChannel(g, c)
	if err != nil {
		fmt.Println("Couldn't set the welcome channel.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		fmt.Println(err.Error())
		return
	}

	// Send feedback
	session.ChannelTyping(c.ID)
	_, err = session.ChannelMessageSend(c.ID, "Le salon de bienvenue est maintenant <#"+c.ID+">.")
	if err != nil {
		fmt.Println("Couldn't announce the new welcome channel.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		fmt.Println(err.Error())
		return
	}
}

func setPresentationChannelCommand(g *discordgo.Guild, c *discordgo.Channel) {

	// Set the presentation channel
	_, err := setPresentationChannel(g, c)
	if err != nil {
		fmt.Println("Couldn't set a presentation channel.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		fmt.Println(err.Error())
		return
	}

	// Send feedback
	session.ChannelTyping(c.ID)
	_, err = session.ChannelMessageSend(c.ID, "Le salon de présentation est maintenant <#"+c.ID+">.")
	if err != nil {
		fmt.Println("Couldn't announce the new presentation channel.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		fmt.Println(err.Error())
		return
	}
}
