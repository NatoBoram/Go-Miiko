package main

import (
	"github.com/bwmarrin/discordgo"
)

func isRole(s *discordgo.Session, g *discordgo.Guild, table string, m *discordgo.Member) (isrole bool, err error) {

	// Get role
	role, err := getRole(s, g, table)
	if err != nil {
		return false, err
	}

	// Compare selected role with every roles
	for _, memberRole := range m.Roles {
		if memberRole == role.ID {
			return true, err
		}
	}

	return false, err
}

func isAdmin(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (isadmin bool, err error) {
	return isRole(s, g, tableAdmin, m)
}

func isMod(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (ismod bool, err error) {
	return isRole(s, g, tableMod, m)
}

func isLight(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (islight bool, err error) {
	return isRole(s, g, tableLight, m)
}

func isAbsynthe(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (isabsynthe bool, err error) {
	return isRole(s, g, tableAbsynthe, m)
}

func isObsidian(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (isobsidian bool, err error) {
	return isRole(s, g, tableObsidian, m)
}

func isShadow(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (isshadow bool, err error) {
	return isRole(s, g, tableShadow, m)
}

func isEel(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (iseel bool, err error) {
	return isRole(s, g, tableEel, m)
}

func isNPC(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (isnpc bool, err error) {
	return isRole(s, g, tableNPC, m)
}
