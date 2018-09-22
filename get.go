package main

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
)

func get(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, ms []string) {
	if len(ms) > 2 {
		switch ms[2] {

		// get lover
		case "lover":
			// GetLoverCmd(db, s, g, c, m.Author)

		// get points
		case "points":
			// GetPoints(s, g, c, m)

		// get presentation
		case "presentation":
			if len(ms) > 3 {
				switch ms[3] {

				// get presentation channel
				case "channel":
					getPresentationChannelCommand(s, g, c)

				// get presentation ?
				default:
					_, err := s.ChannelMessageSend(c.ID, "Erreur sur la commande `"+ms[3]+"`."+"\n"+
						"La commande disponible est `channel`.")
					if err != nil {
						printDiscordError("Couldn't help a get presentation command.", g, c, m, nil, err)
					}
				}
			} else {

				// get presentation
				_, err := s.ChannelMessageSend(c.ID, "La commande disponible est `channel`.")
				if err != nil {
					printDiscordError("Couldn't help a get presentation command.", g, c, m, nil, err)
				}
			}

		// get role
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

				// get role ?
				default:
					_, err := s.ChannelMessageSend(c.ID, "Erreur sur le role `"+ms[3]+"`."+"\n"+
						"Les rôles disponibles sont `admin`, `mod`, `light`, `absynthe`, `obsidian`, `shadow`, `eel` et `npc`.")
					if err != nil {
						printDiscordError("Couldn't help a get role command.", g, c, m, nil, err)
					}
				}
			} else {
				_, err := s.ChannelMessageSend(c.ID, "Les rôles disponibles sont `admin`, `mod`, `light`, `absynthe`, `obsidian`, `shadow`, `eel` et `npc`.")
				if err != nil {
					printDiscordError("Couldn't help a get role command.", g, c, m, nil, err)
				}
			}

		// get roles
		case "roles":
			getRolesCommand(s, g, c)

		// get welcome
		case "welcome":
			if len(ms) > 3 {
				switch ms[3] {

				// get welcome channel
				case "channel":
					getWelcomeChannelCommand(s, g, c)

				// get welcome ?
				default:
					_, err := s.ChannelMessageSend(c.ID, "Erreur sur la commande `"+ms[3]+"`."+"\n"+
						"La commande disponible est `channel`.")
					if err != nil {
						printDiscordError("Couldn't help a get welcome command.", g, c, m, nil, err)
					}
				}
			} else {

				// get welcome
				_, err := s.ChannelMessageSend(c.ID, "La commande disponible est `channel`.")
				if err != nil {
					printDiscordError("Couldn't help a get welcome command.", g, c, m, nil, err)
				}
			}

		// get ?
		default:
			_, err := s.ChannelMessageSend(c.ID, "Erreur sur la commande `"+ms[2]+"`."+"\n"+
				"Les commandes disponibles sont ~~`lover`~~, ~~`points`~~, `presentation`, `role`, `roles` et `welcome`.")
			if err != nil {
				printDiscordError("Couldn't help a set command.", g, c, m, nil, err)
			}
		}

	} else {

		// get
		_, err := s.ChannelMessageSend(c.ID, "Les commandes disponibles sont ~~`lover`~~, ~~`points`~~, `presentation`, `role`, `roles` et `welcome`.")
		if err != nil {
			printDiscordError("Couldn't help a set command.", g, c, m, nil, err)
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
		printDiscordError("Couldn't send the welcome channel.", g, c, nil, nil, err)
		return
	}
}

func getPresentationChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {

	// Get the presentation channel
	channel, err := getPresentationChannel(s, g)
	if err == sql.ErrNoRows {
		_, err = s.ChannelMessageSend(c.ID, "Il n'y a pas de salon de présentation.")
		if err != nil {
			printDiscordError("Couldn't announce the absence of presentation channel.", g, c, nil, nil, err)
		}
		return
	} else if err != nil {
		printDiscordError("Couldn't get the presentation channel.", g, c, nil, nil, err)
		return
	}

	// Send the presentation channel
	s.ChannelMessageSend(c.ID, "Le salon de présentation est <#"+channel.ID+">.")
	if err != nil {
		printDiscordError("Couldn't send the presentation channel.", g, c, nil, nil, err)
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
		printDiscordError("Couldn't send an embed.", g, c, nil, nil, err)
	}
}

func getRoleCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, r *discordgo.Role, err error) {

	// Check Error
	if err == sql.ErrNoRows {
		s.ChannelMessageSend(c.ID, "Je ne connais pas ce rôle.")
		return
	} else if err != nil {
		printDiscordError("Couldn't get a role.", g, c, nil, nil, err)

		_, err = s.ChannelMessageSend(c.ID, "Désolée, je n'ai pas pu trouver ce rôle.")
		if err != nil {
			printDiscordError("Couldn't announce that I couldn't get a role.", g, c, nil, nil, err)
		}

		return
	}

	// Send the role
	s.ChannelTyping(c.ID)
	_, err = s.ChannelMessageSend(c.ID, "Ce rôle est <@&"+r.ID+">.")
	if err != nil {
		printDiscordError("Couldn't tell the role.", g, c, nil, nil, err)
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
