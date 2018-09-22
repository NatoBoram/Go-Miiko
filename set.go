package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Set redirects the `set` command.
func set(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, ms []string) {
	if len(ms) > 2 {
		switch ms[2] {

		// set presentation
		case "presentation":
			if len(ms) > 3 {

				switch ms[3] {

				// set presentation channel
				case "channel":
					setPresentationChannelCommand(s, g, c)

				// set presentation ?
				default:
					_, err := s.ChannelMessageSend(c.ID, "Erreur sur la commande `"+ms[3]+"`."+"\n"+
						"La commande disponible est `channel`.")
					if err != nil {
						printDiscordError("Couldn't help a set presentation command.", g, c, m, nil, err)
					}
				}
			} else {

				// set presentation
				_, err := s.ChannelMessageSend(c.ID, "La commande disponible est `channel`.")
				if err != nil {
					printDiscordError("Couldn't help a set presentation command.", g, c, m, nil, err)
				}
			}

		// set role
		case "role":
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

				// set role ?
				default:
					_, err := s.ChannelMessageSend(c.ID, "Erreur sur le role `"+ms[3]+"`."+"\n"+
						"Les rôles valides sont `admin`, `mod`, `light`, `absynthe`, `obsidian`, `shadow`, `eel` et `npc`.")
					if err != nil {
						printDiscordError("Couldn't help a set role command.", g, c, m, nil, err)
					}
				}
			} else {

				// set role
				_, err := s.ChannelMessageSend(c.ID, "Les rôles disponibles sont `admin`, `mod`, `light`, `absynthe`, `obsidian`, `shadow`, `eel` et `npc`.")
				if err != nil {
					printDiscordError("Couldn't help a set role command.", g, c, m, nil, err)
				}
			}

		// set welcome
		case "welcome":
			if len(ms) > 3 {
				switch ms[3] {

				// set welcome channel
				case "channel":
					setWelcomeChannelCommand(s, g, c)

				// set welcome ?
				default:
					_, err := s.ChannelMessageSend(c.ID, "Erreur sur la commande `"+ms[3]+"`."+"\n"+
						"La commande disponible est `channel`.")
					if err != nil {
						printDiscordError("Couldn't help a set command.", g, c, m, nil, err)
					}
				}
			} else {

				// set welcome
				_, err := s.ChannelMessageSend(c.ID, "La commande disponible est `channel`.")
				if err != nil {
					printDiscordError("Couldn't help a set command.", g, c, m, nil, err)
				}
			}

		// set ?
		default:
			_, err := s.ChannelMessageSend(c.ID, "Erreur sur la commande `"+ms[2]+"`."+"\n"+
				"Les commandes disponibles sont `presentation`, `role` et `welcome`.")
			if err != nil {
				printDiscordError("Couldn't help a set command.", g, c, m, nil, err)
			}
		}
	} else {

		// set
		_, err := s.ChannelMessageSend(c.ID, "Les commandes disponibles sont `presentation`, `role` et `welcome`.")
		if err != nil {
			printDiscordError("Couldn't help a set command.", g, c, m, nil, err)
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
		printDiscordError("Couldn't get a role by its string.", g, c, nil, nil, err)

		_, err := s.ChannelMessageSend(c.ID, "Ce rôle n'existe pas.")
		if err != nil {
			printDiscordError("Couldn't announce that a role doesn't exist.", g, c, nil, nil, err)
		}

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
