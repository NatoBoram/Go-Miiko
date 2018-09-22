package main

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Channels

func getWelcomeChannel(s *discordgo.Session, g *discordgo.Guild) (*discordgo.Channel, error) {

	// Select the presentation channel
	channelID, err := selectWelcomeChannel(g)
	if err != nil {
		return nil, err
	}

	// Turn it into a channel
	return s.Channel(channelID)
}

func getPresentationChannel(s *discordgo.Session, g *discordgo.Guild) (channel *discordgo.Channel, err error) {

	// Select the presentation channel
	channelID, err := selectPresentationChannel(g)
	if err != nil {
		return nil, err
	}

	// Turn it into a channel
	return s.Channel(channelID)
}

// Roles

func getRoles(s *discordgo.Session, g *discordgo.Guild) (admin *discordgo.Role, mod *discordgo.Role, light *discordgo.Role, absynthe *discordgo.Role, obsidian *discordgo.Role, shadow *discordgo.Role, eel *discordgo.Role, npc *discordgo.Role) {

	admin, _ = getRoleAdmin(s, g)
	mod, _ = getRoleMod(s, g)
	light, _ = getRoleLight(s, g)
	absynthe, _ = getRoleAbsynthe(s, g)
	obsidian, _ = getRoleObsidian(s, g)
	shadow, _ = getRoleShadow(s, g)
	eel, _ = getRoleEel(s, g)
	npc, _ = getRoleNPC(s, g)

	return admin, mod, light, absynthe, obsidian, shadow, eel, npc
}

func getRole(s *discordgo.Session, g *discordgo.Guild, table string) (role *discordgo.Role, err error) {

	// Get a role ID from the database
	roleID, err := selectRole(g, table)
	if err != nil {
		return nil, err
	}

	// Get Role
	return s.State.Role(g.ID, roleID)
}

func getRoleAdmin(s *discordgo.Session, g *discordgo.Guild) (role *discordgo.Role, err error) {
	return getRole(s, g, tableAdmin)
}

func getRoleMod(s *discordgo.Session, g *discordgo.Guild) (role *discordgo.Role, err error) {
	return getRole(s, g, tableMod)
}

func getRoleLight(s *discordgo.Session, g *discordgo.Guild) (role *discordgo.Role, err error) {
	return getRole(s, g, tableLight)
}

func getRoleAbsynthe(s *discordgo.Session, g *discordgo.Guild) (role *discordgo.Role, err error) {
	return getRole(s, g, tableAbsynthe)
}

func getRoleObsidian(s *discordgo.Session, g *discordgo.Guild) (role *discordgo.Role, err error) {
	return getRole(s, g, tableObsidian)
}

func getRoleShadow(s *discordgo.Session, g *discordgo.Guild) (role *discordgo.Role, err error) {
	return getRole(s, g, tableShadow)
}

func getRoleEel(s *discordgo.Session, g *discordgo.Guild) (role *discordgo.Role, err error) {
	return getRole(s, g, tableEel)
}

func getRoleNPC(s *discordgo.Session, g *discordgo.Guild) (role *discordgo.Role, err error) {
	return getRole(s, g, tableNPC)
}

func getRoleByString(s *discordgo.Session, g *discordgo.Guild, str string) (role *discordgo.Role, err error) {

	// Get roles
	roles, err := s.GuildRoles(g.ID)
	if err != nil {
		fmt.Println("Couldn't get the guild's roles.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Role :", str)
		fmt.Println(err.Error())
		return
	}

	// Return the first occurence
	for _, role := range roles {
		if role.Name == str || role.ID == str {
			return role, nil
		}
	}

	return nil, errors.New("this role doesn't exist")
}
