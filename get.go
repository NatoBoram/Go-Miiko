package main

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func get(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, ms []string) {
	if len(ms) > 2 {
		switch ms[2] {
		case "welcome":
			if len(ms) > 3 {
				if ms[3] == "channel" {
					getWelcomeChannelCommand(s, g, c)
				}
			}
		case "presentation":
			if len(ms) > 3 {
				if ms[3] == "channel" {
					getPresentationChannelCommand(s, g, c)
				}
			}
		case "points":
			// GetPoints(s, g, c, m)
		case "roles":
			getRolesCommand(s, g, c)
		case "role":
			if len(ms) > 3 {
				switch ms[3] {
				case "admin":
					getRoleAdminCommand(s, g, c)
				case "mod":
					getRoleModCommand(s, g, c)
				case "light":
					getRoleLightCommand(s, g, c)
				case "absynthe":
					getRoleAbsyntheCommand(s, g, c)
				case "obsidian":
					getRoleObsidianCommand(s, g, c)
				case "shadow":
					getRoleShadowCommand(s, g, c)
				case "eel":
					getRoleEelCommand(s, g, c)
				case "npc":
					getRoleNPCCommand(s, g, c)
				}
			}
		case "lover":
			// GetLoverCmd(db, s, g, c, m.Author)
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

func newRoleEmbedField(name string, r *discordgo.Role) *discordgo.MessageEmbedField {
	value := "­" // Zero-Width Space
	if r != nil {
		value = r.Name
	}
	return &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: true,
	}
}

func getRolesCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {

	// Get Roles
	admin, mod, light, absynthe, obsidian, shadow, eel, npc := getRoles(s, g)

	// Create Embed
	embed := &discordgo.MessageEmbed{
		Title:       "Rôles",
		Color:       0x34386f,
		Description: "Voici les rôles que je connais dans **" + g.Name + "**.",
		Fields: []*discordgo.MessageEmbedField{
			newRoleEmbedField("Administration", admin),
			newRoleEmbedField("Modération", mod),
			newRoleEmbedField("Étincelante", light),
			newRoleEmbedField("Absynthe", absynthe),
			newRoleEmbedField("Obsidienne", obsidian),
			newRoleEmbedField("Ombre", shadow),
			newRoleEmbedField("Eel", eel),
			newRoleEmbedField("PNJ", npc),
		},
	}

	s.ChannelTyping(c.ID)
	_, err := s.ChannelMessageSendEmbed(c.ID, embed)
	if err != nil {
		fmt.Println("Couldn't send an embed.")
		fmt.Println(err.Error())
	}
}

func getRoleCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, r *discordgo.Role, err error) {

	// Check Error
	if err == sql.ErrNoRows {
		s.ChannelMessageSend(c.ID, "Je ne connais pas ce rôle.")
		return
	} else if err != nil {
		s.ChannelMessageSend(c.ID, "Désolée, je n'ai pas pu trouver ce rôle.")
		fmt.Println(err.Error())
		return
	}

	// Send the role
	s.ChannelTyping(c.ID)
	_, err = s.ChannelMessageSend(c.ID, "Ce rôle est <@&"+r.ID+">.")
	if err != nil {
		fmt.Println("Couldn't tell the role.")
		fmt.Println(err.Error())
	}
}

func getRoleAdminCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {
	role, err := getRoleAdmin(s, g)
	getRoleCommand(s, g, c, role, err)
}

func getRoleModCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {
	role, err := getRoleMod(s, g)
	getRoleCommand(s, g, c, role, err)
}

func getRoleLightCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {
	role, err := getRoleLight(s, g)
	getRoleCommand(s, g, c, role, err)
}

func getRoleAbsyntheCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {
	role, err := getRoleAbsynthe(s, g)
	getRoleCommand(s, g, c, role, err)
}

func getRoleObsidianCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {
	role, err := getRoleObsidian(s, g)
	getRoleCommand(s, g, c, role, err)
}

func getRoleShadowCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {
	role, err := getRoleShadow(s, g)
	getRoleCommand(s, g, c, role, err)
}

func getRoleEelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {
	role, err := getRoleEel(s, g)
	getRoleCommand(s, g, c, role, err)
}

func getRoleNPCCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {
	role, err := getRoleNPC(s, g)
	getRoleCommand(s, g, c, role, err)
}
