package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Set redirects the `set` coommand.
func set(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, ms []string) {
	if len(ms) > 2 {
		switch ms[2] {
		case "welcome":
			if len(ms) > 3 {
				if ms[3] == "channel" {
					if m.Author.ID == g.OwnerID {
						setWelcomeChannelCommand(s, g, c)
					}
				}
			}
		case "presentation":
			if len(ms) > 3 {
				if ms[3] == "channel" {
					if m.Author.ID == g.OwnerID {
						setPresentationChannelCommand(s, g, c)
					}
				}
			}
		case "role":
			if m.Author.ID != g.OwnerID {
				s.ChannelMessageSend(c.ID, "Désolée, mais vous n'avez pas les permissions nécessaires.")
				return
			}
			if len(ms) > 4 {
				switch ms[3] {
				case "admin":
					setRoleAdminCommand(s, g, c, ms[4])
				case "mod":
					setRoleModCommand(s, g, c, ms[4])
				case "light":
					setRoleLightCommand(s, g, c, ms[4])
				case "absynthe":
					setRoleAbsyntheCommand(s, g, c, ms[4])
				case "obsidian":
					setRoleObsidianCommand(s, g, c, ms[4])
				case "shadow":
					setRoleShadowCommand(s, g, c, ms[4])
				case "eel":
					setRoleEelCommand(s, g, c, ms[4])
				case "npc":
					setRoleNPCCommand(s, g, c, ms[4])
				default:
					s.ChannelMessageSend(c.ID, "Les rôles valides sont `admin`, `mod`, `light`, `absynthe`, `obsidian`, `shadow`, `eel` et `npc`.")
				}
			} else {

			}
		}
	}
}

// setWelcomeChannelCommand sets the welcome channel and sends feedback to the user.
func setWelcomeChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {

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
	s.ChannelTyping(c.ID)
	_, err = s.ChannelMessageSend(c.ID, "Le salon de bienvenue est maintenant <#"+c.ID+">.")
	if err != nil {
		fmt.Println("Couldn't announce the new welcome channel.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		fmt.Println(err.Error())
		return
	}
}

func setPresentationChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {
	s.ChannelTyping(c.ID)

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
	s.ChannelTyping(c.ID)
	_, err = s.ChannelMessageSend(c.ID, "Le salon de présentation est maintenant <#"+c.ID+">.")
	if err != nil {
		fmt.Println("Couldn't announce the new presentation channel.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		fmt.Println(err.Error())
		return
	}
}

func setRoleCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, roleString string, table string) {
	s.ChannelTyping(c.ID)

	// Get a role from the command
	role, err := getRoleByString(s, g, roleString)
	if err != nil {
		fmt.Println("Couldn't get a role by its string.")
		fmt.Println(err.Error())
		s.ChannelMessageSend(c.ID, "Ce rôle n'existe pas.")
		return
	}

	// Set the role
	res, err := setRole(s, g, role, table)
	if err != nil {
		fmt.Println("Couldn't set a role.")
		fmt.Println(res)
		fmt.Println(err.Error())
		s.ChannelMessageSend(c.ID, "Désolée, je n'ai pas pu sauvegarder ce rôle.")
		return
	}

	// Announce the new role
	_, err = s.ChannelMessageSend(c.ID, "Ce role est maintenant <@&"+role.ID+">.")
	if err != nil {
		fmt.Println("Couldn't announce the new role.")
		fmt.Println(err.Error())
	}
}

func setRoleAdminCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, roleString string) {
	setRoleCommand(s, g, c, roleString, tableAdmin)
}

func setRoleModCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, roleString string) {
	setRoleCommand(s, g, c, roleString, tableMod)
}

func setRoleLightCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, roleString string) {
	setRoleCommand(s, g, c, roleString, tableLight)
}

func setRoleAbsyntheCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, roleString string) {
	setRoleCommand(s, g, c, roleString, tableAbsynthe)
}

func setRoleObsidianCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, roleString string) {
	setRoleCommand(s, g, c, roleString, tableObsidian)
}

func setRoleShadowCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, roleString string) {
	setRoleCommand(s, g, c, roleString, tableShadow)
}

func setRoleEelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, roleString string) {
	setRoleCommand(s, g, c, roleString, tableEel)
}

func setRoleNPCCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, roleString string) {
	setRoleCommand(s, g, c, roleString, tableNPC)
}
