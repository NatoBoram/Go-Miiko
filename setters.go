package main

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
)

func setPresentationChannel(g *discordgo.Guild, c *discordgo.Channel) (sql.Result, error) {

	// Check if there's a presentation channel
	_, err := selectPresentationChannel(g)
	if err == sql.ErrNoRows {

		// Insert if there's none
		return insertPresentationChannel(g, c)
	} else if err != nil {
		return nil, err
	}

	// Update if there's one
	return updatePresentationChannel(g, c)
}

func setWelcomeChannel(g *discordgo.Guild, c *discordgo.Channel) (sql.Result, error) {

	// Check if there's a presentation channel
	_, err := selectWelcomeChannel(g)
	if err == sql.ErrNoRows {

		// Insert if there's none
		return insertWelcomeChannel(g, c)
	} else if err != nil {
		return nil, err
	}

	// Update if there's one
	return updateWelcomeChannel(g, c)
}

func setRole(s *discordgo.Session, g *discordgo.Guild, r *discordgo.Role, table string) (res sql.Result, err error) {

	// Check if the role exists
	_, err = selectRole(g, table)
	if err == sql.ErrNoRows {

		// Insert if there's none
		return insertRole(g, r, table)
	} else if err != nil {
		return nil, err
	}

	// Update if there's one
	return updateRole(g, r, table)
}

func setRoleAdmin(s *discordgo.Session, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(s, g, r, tableAdmin)
}

func setRoleMod(s *discordgo.Session, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(s, g, r, tableMod)
}

func setRoleLight(s *discordgo.Session, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(s, g, r, tableLight)
}

func setRoleAbsynthe(s *discordgo.Session, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(s, g, r, tableAbsynthe)
}

func setRoleObsidian(s *discordgo.Session, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(s, g, r, tableObsidian)
}

func setRoleShadow(s *discordgo.Session, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(s, g, r, tableShadow)
}

func setRoleEel(s *discordgo.Session, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(s, g, r, tableEel)
}

func setRoleNPC(s *discordgo.Session, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(s, g, r, tableNPC)
}
