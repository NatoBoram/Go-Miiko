package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/NatoBoram/Go-Miiko/wheel"
)

// PlaceInAGuard gives members a role.
func placeInAGuard(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, u *discordgo.Member, m *discordgo.Message) bool {

	// No guard for bots!
	if m.Author.Bot {
		return false
	}

	// Check if author has no role
	has, err := hasGuard(s, g, u)
	if err != nil {
		printDiscordError("Couldn't check if a member has a guard.", g, c, m, nil, err)
		return false
	} else if has {
		return false
	}

	// Get mentionned roles
	guards, tables, err := getMentionnedGuard(s, g, m)

	// Check if there's only one mentionned role
	var guard *discordgo.Role
	var table string
	if len(guards) == 1 {
		guard = guards[0]
		table = tables[0]
	} else if len(guards) > 1 {

		// If there's more than one...
		s.ChannelMessageSend(c.ID, "Désolée, mais je ne peux t'offrir qu'une seule garde.")
		return true
	} else if len(guards) == 0 {

		// Why are you ignoring me?
		if wheel.RandomOverPhiPower(3) {

			// Typing!
			err = s.ChannelTyping(m.ChannelID)
			if err != nil {
				printDiscordError("Couldn't tell that I'm typing.", g, c, m, nil, err)
			}

			// Protest
			_, err = s.ChannelMessageSend(m.ChannelID, getProtestMessage(u.User))
			if err != nil {
				printDiscordError("Couldn't protest being ignored.", g, c, m, nil, err)
				return false
			}
			return true
		}
		return false
	}

	// Typing!
	err = s.ChannelTyping(m.ChannelID)
	if err != nil {
		printDiscordError("Couldn't tell that I'm typing.", g, c, m, nil, err)
	}

	if table == tableLight {
		_, err := s.ChannelMessageSend(m.ChannelID, "Si tu fais partie de la Garde <@&"+guard.ID+">, envoie un message à <@"+g.OwnerID+"> sur Eldarya pour annoncer ta présence. En attendant, dans quelle garde est ton personnage sur Eldarya?")
		if err != nil {
			printDiscordError("Couldn't send message for special role.", g, c, m, nil, err)
		}
	} else if table == tableAbsynthe || table == tableObsidian || table == tableShadow {

		// Add role
		err := s.GuildMemberRoleAdd(g.ID, u.User.ID, guard.ID)
		if err != nil {
			printDiscordError("Couldn't add a role.", g, c, m, nil, err)
			_, err = s.ChannelMessageSend(m.ChannelID, "<@"+g.OwnerID+"> Je n'ai pas pu donner le rôle <@&"+guard.ID+"> à <@"+u.User.ID+">. Peut-être que je n'ai pas la permission?")
			return false
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, getGuardMessage(u.User, guard))
		if err != nil {
			printDiscordError("Couldn't announce new role.", g, c, m, nil, err)
		}

		// Once a valid guard is received, ask to introduce in the appropriate channel.
		err = askForIntroduction(s, g, c)
		if err != nil {
			printDiscordError("Couldn't announce the introduction channel.", g, c, m, nil, err)
		}

	} else if table == tableEel {

		// Add role
		err := s.GuildMemberRoleAdd(g.ID, u.User.ID, guard.ID)
		if err != nil {
			printDiscordError("Couldn't add a role.", g, c, m, nil, err)
			_, err = s.ChannelMessageSend(m.ChannelID, "<@"+g.OwnerID+"> Je n'ai pas pu donner le rôle <@&"+guard.ID+"> à <@"+u.User.ID+">. Peut-être que je n'ai pas la permission?")
			return false
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, "D'accord, <@"+u.User.ID+">. Je t'ai donné le rôle <@&"+guard.ID+"> en attendant que tu rejoignes une garde.")
		if err != nil {
			printDiscordError("Couldn't announce new role.", g, c, m, nil, err)
		}

	} else if table == tableNPC {

		// Add role
		err := s.GuildMemberRoleAdd(g.ID, u.User.ID, guard.ID)
		if err != nil {
			printDiscordError("Couldn't add a role.", g, c, m, nil, err)
			_, err = s.ChannelMessageSend(m.ChannelID, "<@"+g.OwnerID+"> Je n'ai pas pu donner le rôle <@&"+guard.ID+"> à <@"+u.User.ID+">. Peut-être que je n'ai pas la permission?")
			return false
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, "D'accord, <@"+u.User.ID+">. Je t'ai donné le rôle <@&"+guard.ID+">, mais saches que ce serveur est dédié à Eldarya.")
		if err != nil {
			printDiscordError("Couldn't announce new role.", g, c, m, nil, err)
		}
	}

	// Status
	err = setStatus(s, "accueillir "+u.User.Username+" dans "+g.Name)
	if err != nil {
		printDiscordError("Couldn't set the status to welcoming someone.", g, c, m, nil, err)
	}

	return true
}

func getMentionnedGuard(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Message) (guards []*discordgo.Role, tables []string, err error) {

	// Light
	if strings.Contains(strings.ToLower(m.Content), "tincelant") {
		role, err := getRoleLight(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Light guard.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableLight)
		}
	}

	// Obsidian
	if strings.Contains(strings.ToLower(m.Content), "obsi") {
		role, err := getRoleObsidian(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Obsidian guard.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableObsidian)
		}
	}

	// Absynthe
	if strings.Contains(strings.ToLower(m.Content), "absy") {
		role, err := getRoleAbsynthe(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Absynthe guard.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableAbsynthe)
		}
	}

	// Shadow
	if strings.Contains(strings.ToLower(m.Content), "ombr") {
		role, err := getRoleShadow(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Shadow guard.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableShadow)
		}
	}

	// Eel
	if strings.Contains(strings.ToLower(m.Content), "eel") || strings.Contains(strings.ToLower(m.Content), "aucun") || strings.Contains(strings.ToLower(m.Content), "ai pas") || strings.Contains(strings.ToLower(m.Content), "pas encore") || strings.Contains(strings.ToLower(m.Content), "de commencer") {
		role, err := getRoleEel(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Eel guard.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableEel)
		}
	}

	// NPC
	if strings.Contains(strings.ToLower(m.Content), "joue pas") || strings.Contains(strings.ToLower(m.Content), " quoi") || strings.Contains(strings.ToLower(m.Content), "pas commencé") {
		role, err := getRoleNPC(s, g)
		if err != nil {
			fmt.Println("Couldn't get the NPC role.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableNPC)
		}
	}
	return
}

func createGuard(s *discordgo.Session, g *discordgo.Guild, name string) *discordgo.Role {

	// Get color
	var color int
	if name == "Étincelante" {
		color = 16705182
	} else if name == "Obsidienne" {
		color = 16496296
	} else if name == "Absynthe" {
		color = 8321915
	} else if name == "Ombre" {
		color = 16364540
	} else if name == "Eel" {
		color = 12503544
	} else if name == "PNJ" {
		color = 10263708
	} else {
		return nil
	}

	// Create the missing role
	role, err := s.GuildRoleCreate(g.ID)
	if err != nil {
		fmt.Println("Couldn't create the role", name, "in", g.Name+".")
		fmt.Println(err.Error())
		return nil
	}

	// Edit the missing role
	_, err = s.GuildRoleEdit(g.ID, role.ID, name, color, false, role.Permissions, false)
	if err != nil {
		fmt.Println("Couldn't edit the newly created role,", name+".")
		fmt.Println(err.Error())
		return nil
	}

	return role
}

func getGuardMessage(user *discordgo.User, role *discordgo.Role) string {

	// Messages
	messageList := [...]string{
		"Bienvenue à <@" + user.ID + "> dans la garde <@&" + role.ID + ">!",
		"Bienvenue dans la garde <@&" + role.ID + ">, <@" + user.ID + ">.",
		"Bienvenue dans la garde <@&" + role.ID + ">, <@" + user.ID + ">. J'espère que tu ne t'attends pas à un matelas.",
		"Bienvenue au sein la garde <@&" + role.ID + ">, <@" + user.ID + ">.",
		"Bienvenue parmis les <@&" + role.ID + ">, <@" + user.ID + ">.",
		"<@" + user.ID + "> est maintenant un membre de la garde <@&" + role.ID + ">!",
		"<@&" + role.ID + "> a l'honneur d'accueillir <@" + user.ID + ">!",
		"<@&" + role.ID + ">! Faites de la place pour <@" + user.ID + ">!",
		"<@" + user.ID + "> fait maintenant partie de la garde <@&" + role.ID + ">.",
		"<@" + user.ID + "> est maintenant une <@&" + role.ID + ">!",
		"Souhaitez la bienvenue à notre nouvelle <@&" + role.ID + ">, <@" + user.ID + ">!",
		"Bien! <@" + user.ID + "> a maintenant une place dans les cachots de la garde <@&" + role.ID + ">.",
		"<@" + user.ID + "> a rejoint la garde <@&" + role.ID + ">.",
		"Je savais que <@" + user.ID + "> était une <@&" + role.ID + ">!",
		"Ah, je savais que <@" + user.ID + "> était une <@&" + role.ID + ">.",
		"Je savais bien que <@" + user.ID + "> était une <@&" + role.ID + ">!",
		"Ah, je le savais! <@" + user.ID + "> est une <@&" + role.ID + ">!",
		"J'en étais sûre! <@" + user.ID + "> est une <@&" + role.ID + ">!",
		"<@" + user.ID + "> est dorénavant une <@&" + role.ID + ">.",
		"Accueillez notre nouvelle <@&" + role.ID + ">, <@" + user.ID + ">!",
		"Je te souhaite un bon séjour parmis les <@&" + role.ID + ">, <@" + user.ID + ">.",
		"<@" + user.ID + "> peut maintenant rejoindre les <@&" + role.ID + ">.",
		"Tu peux rejoindre les <@&" + role.ID + ">, <@" + user.ID + ">.",
		"Que les <@&" + role.ID + "> soient avec <@" + user.ID + ">.",
	}

	// Seed
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return messageList[random.Intn(len(messageList))]
}

func getProtestMessage(u *discordgo.User) string {

	// Messages
	protestList := [...]string{
		"J'apprécierais de ne pas me faire ignorer, <@" + u.ID + ">.",
		"J'apprécierais vraiment de ne pas me faire ignorer, <@" + u.ID + ">.",
		"J'apprécierais *vraiment* de ne pas me faire ignorer, <@" + u.ID + ">.",

		"<@" + u.ID + ">, j'apprécierais de ne pas me faire ignorer.",
		"<@" + u.ID + ">, j'apprécierais vraiment de ne pas me faire ignorer.",
		"<@" + u.ID + ">, j'apprécierais *vraiment* de ne pas me faire ignorer.",

		"Je déteste me faire ignorer, <@" + u.ID + ">.",
		"Je déteste vraiment me faire ignorer, <@" + u.ID + ">.",
		"Je déteste *vraiment* me faire ignorer, <@" + u.ID + ">.",

		"<@" + u.ID + ">, je déteste me faire ignorer.",
		"<@" + u.ID + ">, je déteste vraiment me faire ignorer.",
		"<@" + u.ID + ">, je déteste *vraiment* me faire ignorer.",

		"Tu dois m'indiquer ta garde, <@" + u.ID + ">.",
		"<@" + u.ID + ">, tu dois m'indiquer ta garde, .",

		"Je me répète, <@" + u.ID + ">.",
		"<@" + u.ID + ">, je me répète :",
	}

	// Ask for guard, but less friendly
	askList := [...]string{
		"Quelle est ta garde?",
		"De quelle garde fais-tu partie?",
		"Dans quelle garde es-tu?",
	}

	// Seed
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return protestList[random.Intn(len(protestList))] + " " + askList[random.Intn(len(askList))]
}
