package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Set redirects the `set` command.
func set(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, ms []string) {
	if len(ms) > 2 {
		switch ms[2] {

		// set fame
		case "fame":
			if len(ms) > 3 {
				switch ms[3] {

				// set fame channel
				case "channel":
					setFameChannelCommand(s, g, c, m)

				// set fame ?
				default:
					_, err := s.ChannelMessageSend(c.ID, "Erreur sur la commande `"+ms[3]+"`."+"\n"+
						"La commande disponible est `channel`.")
					if err != nil {
						printDiscordError("Couldn't help a set fame command.", g, c, m, nil, err)
					}
				}
			} else {

				// set fame
				_, err := s.ChannelMessageSend(c.ID, "La commande disponible est `channel`.")
				if err != nil {
					printDiscordError("Couldn't help a set fame command.", g, c, m, nil, err)
				}
			}

		// set presentation
		case "presentation":
			if len(ms) > 3 {

				switch ms[3] {

				// set presentation channel
				case "channel":
					setPresentationChannelCommand(s, g, c, m)

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
					setRoleAdminCommand(s, g, c, m, strings.Join(ms[4:], " "))
				case "mod":
					setRoleModCommand(s, g, c, m, strings.Join(ms[4:], " "))
				case "light":
					setRoleLightCommand(s, g, c, m, strings.Join(ms[4:], " "))
				case "absynthe":
					setRoleAbsyntheCommand(s, g, c, m, strings.Join(ms[4:], " "))
				case "obsidian":
					setRoleObsidianCommand(s, g, c, m, strings.Join(ms[4:], " "))
				case "shadow":
					setRoleShadowCommand(s, g, c, m, strings.Join(ms[4:], " "))
				case "eel":
					setRoleEelCommand(s, g, c, m, strings.Join(ms[4:], " "))
				case "npc":
					setRoleNPCCommand(s, g, c, m, strings.Join(ms[4:], " "))

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

		// set sar
		case "sar":
			if len(ms) > 3 {

				switch ms[3] {

				// set sar add
				case "add":
					if len(ms) > 4 {
						setSARAddCommand(s, g, c, m, strings.Join(ms[4:], " "))
					} else {
						_, err := s.ChannelMessageSend(c.ID, "Vous devez spécifier un rôle.")
						if err != nil {
							printDiscordError("Couldn't help a set sar add command.", g, c, m, nil, err)
						}
					}

				// set sar remove
				case "remove":
					if len(ms) > 4 {
						setSARRemoveCommand(s, g, c, m, strings.Join(ms[4:], " "))
					} else {
						_, err := s.ChannelMessageSend(c.ID, "Vous devez spécifier un rôle.")
						if err != nil {
							printDiscordError("Couldn't help a set sar remove command.", g, c, m, nil, err)
						}
					}

				// set sar ?
				default:
					_, err := s.ChannelMessageSend(c.ID, "Erreur sur la commande `"+ms[3]+"`."+"\n"+
						"Les commandes disponibles sont `add` et `remove`.")
					if err != nil {
						printDiscordError("Couldn't help a set sar command.", g, c, m, nil, err)
					}
				}
			} else {

				// set sar
				_, err := s.ChannelMessageSend(c.ID, "Les commandes disponibles sont `add` et `remove`.")
				if err != nil {
					printDiscordError("Couldn't help a set sar command.", g, c, m, nil, err)
				}
			}

		// set status
		case "status":

			// Lock to master only
			if m.Author.ID != master.ID {
				_, err := s.ChannelMessageSend(c.ID, "Cette commande est réservée à **"+master.Username+"**.")
				if err != nil {
					printDiscordError("Couldn't tell that set status is locked to owner.", g, c, m, nil, err)
				}
				return
			}

			if len(ms) > 3 {

				// Update status
				err := setStatus(s, strings.Join(ms[3:], " "))
				if err != nil {
					printDiscordError("Couldn't set the status to "+strings.Join(ms[3:], " ")+".", g, c, m, nil, err)
				}

			} else {

				// set status
				err := setStatus(s, "")
				if err != nil {
					printDiscordError("Couldn't clear the status at "+master.Username+"'s demand.", g, c, m, nil, err)
				}
			}

		// set welcome
		case "welcome":
			if len(ms) > 3 {
				switch ms[3] {

				// set welcome channel
				case "channel":
					setWelcomeChannelCommand(s, g, c, m)

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
		_, err := s.ChannelMessageSend(c.ID, "Les commandes disponibles sont `fame`, `presentation`, `role`, `sar` et `welcome`.")
		if err != nil {
			printDiscordError("Couldn't help a set command.", g, c, m, nil, err)
		}
	}
}

func setWelcomeChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message) {
	s.ChannelTyping(c.ID)

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
	_, err = s.ChannelMessageSend(c.ID, "Le salon de bienvenue est maintenant <#"+c.ID+">.")
	if err != nil {
		fmt.Println("Couldn't announce the new welcome channel.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		fmt.Println(err.Error())
		return
	}
}

func setPresentationChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message) {
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
	_, err = s.ChannelMessageSend(c.ID, "Le salon de présentation est maintenant <#"+c.ID+">.")
	if err != nil {
		fmt.Println("Couldn't announce the new presentation channel.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		fmt.Println(err.Error())
		return
	}
}

func setFameChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message) {
	s.ChannelTyping(c.ID)

	// Set the welcome channel
	_, err := setFameChannel(g, c)
	if err != nil {
		printDiscordError("Couldn't set the hall of fame channel.", g, c, m, nil, err)
		return
	}

	// Send feedback
	_, err = s.ChannelMessageSend(c.ID, "Le salon de renommée est maintenant <#"+c.ID+">.")
	if err != nil {
		printDiscordError("Couldn't announce the new hall of fame channel.", g, c, m, nil, err)
		return
	}
}

func setRoleCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string, table string) {
	s.ChannelTyping(c.ID)

	// Get a role from the command
	role, err := getRoleByString(s, g, roleString)
	if err != nil {
		printDiscordError("Couldn't get a role by its string.", g, c, m, nil, err)

		_, err := s.ChannelMessageSend(c.ID, "Ce rôle n'existe pas.")
		if err != nil {
			printDiscordError("Couldn't announce that a role doesn't exist.", g, c, m, nil, err)
		}

		return
	}

	// Set the role
	_, err = setRole(g, role, table)
	if err != nil {
		printDiscordError("Couldn't set a role.", g, c, m, nil, err)

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

func setRoleAdminCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	setRoleCommand(s, g, c, m, roleString, tableAdmin)
}

func setRoleModCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	setRoleCommand(s, g, c, m, roleString, tableMod)
}

func setRoleLightCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	setRoleCommand(s, g, c, m, roleString, tableLight)
}

func setRoleAbsyntheCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	setRoleCommand(s, g, c, m, roleString, tableAbsynthe)
}

func setRoleObsidianCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	setRoleCommand(s, g, c, m, roleString, tableObsidian)
}

func setRoleShadowCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	setRoleCommand(s, g, c, m, roleString, tableShadow)
}

func setRoleEelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	setRoleCommand(s, g, c, m, roleString, tableEel)
}

func setRoleNPCCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	setRoleCommand(s, g, c, m, roleString, tableNPC)
}

// Self-Assignable Roles

func setSARAddCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	s.ChannelTyping(c.ID)

	// Get a role from the command
	role, err := getRoleByString(s, g, roleString)
	if err != nil {
		printDiscordError("Couldn't get a role by its string.", g, c, m, nil, err)

		// Announce error
		_, err := s.ChannelMessageSend(c.ID, "Ce rôle n'existe pas.")
		if err != nil {
			printDiscordError("Couldn't announce that a role doesn't exist.", g, c, m, nil, err)
		}
		return
	}

	// Set the role
	_, err = setSAR(g, role)
	if err != nil {
		printDiscordError("Couldn't set a self-assigned role.", g, c, m, nil, err)

		// Announce error
		_, err = s.ChannelMessageSend(c.ID, "Désolée, je n'ai pas pu sauvegarder ce rôle.")
		if err != nil {
			printDiscordError("Couldn't announce that I couldn't save a role.", g, c, m, nil, err)
		}
		return
	}

	// Announce the new role
	_, err = s.ChannelMessageSend(c.ID, "Le rôle <@&"+role.ID+"> est maintenant auto-assignable.")
	if err != nil {
		printDiscordError("Couldn't announce the new SAR.", g, c, m, nil, err)
	}
}

func setSARRemoveCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, roleString string) {
	s.ChannelTyping(c.ID)

	// Get a role from the command
	role, err := getRoleByString(s, g, roleString)
	if err != nil {
		printDiscordError("Couldn't get a role by its string.", g, c, m, nil, err)

		// Announce error
		_, err := s.ChannelMessageSend(c.ID, "Ce rôle n'existe pas.")
		if err != nil {
			printDiscordError("Couldn't announce that a role doesn't exist.", g, c, m, nil, err)
		}
		return
	}

	// Set the role
	_, err = deleteSAR(g, role)
	if err != nil {
		printDiscordError("Couldn't delete a self-assigned role.", g, c, m, nil, err)

		// Announce error
		_, err = s.ChannelMessageSend(c.ID, "Désolée, je n'ai pas pu retirer ce rôle.")
		if err != nil {
			printDiscordError("Couldn't announce that I couldn't remove a role.", g, c, m, nil, err)
		}
		return
	}

	// Announce the new role
	_, err = s.ChannelMessageSend(c.ID, "Le rôle <@&"+role.ID+"> n'est plus auto-assignable.")
	if err != nil {
		printDiscordError("Couldn't announce a removed SAR.", g, c, m, nil, err)
	}
}
