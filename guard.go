package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// PlaceInAGuard gives members a role.
func placeInAGuard(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, u *discordgo.Member, m *discordgo.Message) bool {

	// If Author has no role
	if len(u.Roles) != 0 || m.Author.Bot {
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
		s.ChannelMessageSend(c.ID, "Désolée, mais je ne peux t'offrir qu'une seule garde.")
		return false
	} else {
		// Why are you ignoring me?
		return false
	}

	// Typing!
	err = s.ChannelTyping(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't tell that I'm typing.")
		fmt.Println("Channel : " + c.Name)
		fmt.Println(err.Error())
	}

	if table == tableLight {
		_, err := s.ChannelMessageSend(m.ChannelID, "Si tu fais partie de la Garde <@&"+guard.ID+">, envoie un message à <@"+g.OwnerID+"> sur Eldarya pour annoncer ta présence. En attendant, dans quelle garde est ton personnage sur Eldarya?")
		if err != nil {
			fmt.Println("Couldn't send message for special role.")
			fmt.Println("Channel : " + c.Name)
			fmt.Println(err.Error())
		}
	} else if table == tableAbsynthe || table == tableObsidian || table == tableShadow {

		// Add role
		err := s.GuildMemberRoleAdd(g.ID, u.User.ID, guard.ID)
		if err != nil {
			fmt.Println("Couldn't add a role.")
			fmt.Println("Guild : " + g.Name)
			fmt.Println("Role : " + guard.ID)
			fmt.Println("Member : " + u.User.Username)
			fmt.Println(err.Error())
			return false
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, getGuardMessage(u.User, guard))
		if err != nil {
			fmt.Println("Couldn't announce new role.")
			fmt.Println("Channel : " + c.Name)
			fmt.Println(err.Error())
		}

		// Once a valid guard is received, ask to introduce in the appropriate channel.
		askForIntroduction(s, g, c)

	} else if table == tableEel {

		// Add role
		err := s.GuildMemberRoleAdd(g.ID, u.User.ID, guard.ID)
		if err != nil {
			fmt.Println("Couldn't add a role.")
			fmt.Println("Guild : " + g.Name)
			fmt.Println("Member : " + m.Author.Username)
			fmt.Println(err.Error())
			return false
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, "D'accord, <@"+u.User.ID+">. Je t'ai donné le rôle <@&"+guard.ID+"> en attendant que tu rejoignes une garde.")
		if err != nil {
			fmt.Println("Couldn't announce new role.")
			fmt.Println("Channel : " + c.Name)
			fmt.Println(err.Error())
		}

	} else if table == tableNPC {

		// Add role
		err := s.GuildMemberRoleAdd(g.ID, u.User.ID, guard.ID)
		if err != nil {
			fmt.Println("Couldn't add a role.")
			fmt.Println("Guild : " + g.Name)
			fmt.Println("Member : " + m.Author.Username)
			fmt.Println(err.Error())
			return false
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, "D'accord, <@"+u.User.ID+">. Je t'ai donné le rôle <@&"+guard.ID+">, mais saches que ce serveur est dédié à Eldarya.")
		if err != nil {
			fmt.Println("Couldn't announce new role.")
			fmt.Println("Channel : " + c.Name)
			fmt.Println(err.Error())
		}
	}

	return true
}

func getMentionnedGuard(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Message) (guards []*discordgo.Role, tables []string, err error) {
	if strings.Contains(strings.ToLower(m.Content), "tincelant") {
		role, err := getRoleLight(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Light guard.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableLight)
		}
	}
	if strings.Contains(strings.ToLower(m.Content), "obsi") {
		role, err := getRoleObsidian(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Light guard.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableObsidian)
		}
	}
	if strings.Contains(strings.ToLower(m.Content), "absy") {
		role, err := getRoleAbsynthe(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Light guard.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableAbsynthe)
		}
	}
	if strings.Contains(strings.ToLower(m.Content), "ombr") {
		role, err := getRoleShadow(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Light guard.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableShadow)
		}
	}
	if strings.Contains(strings.ToLower(m.Content), "eel") || strings.Contains(strings.ToLower(m.Content), "aucun") || strings.Contains(strings.ToLower(m.Content), "ai pas") || strings.Contains(strings.ToLower(m.Content), "pas encore") || strings.Contains(strings.ToLower(m.Content), "de commencer") {
		role, err := getRoleEel(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Light guard.")
		} else {
			guards = append(guards, role)
			tables = append(tables, tableEel)
		}
	}
	if strings.Contains(strings.ToLower(m.Content), "joue pas") || strings.Contains(strings.ToLower(m.Content), " quoi") || strings.Contains(strings.ToLower(m.Content), "pas commencé") {
		role, err := getRoleNPC(s, g)
		if err != nil {
			fmt.Println("Couldn't get the Light guard.")
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
