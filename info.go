package main

import "github.com/bwmarrin/discordgo"

func info(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, ms []string) {

	// info
	if len(ms) > 2 {
		switch ms[2] {
		case "channel":
		case "member":
		case "user":

		// info ?
		default:
		}
	} else {
		// info

	}
}
