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

func setRole(g *discordgo.Guild, r *discordgo.Role, table string) (res sql.Result, err error) {

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

func setRoleAdmin(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(g, r, tableAdmin)
}

func setRoleMod(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(g, r, tableMod)
}

func setRoleLight(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(g, r, tableLight)
}

func setRoleAbsynthe(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(g, r, tableAbsynthe)
}

func setRoleObsidian(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(g, r, tableObsidian)
}

func setRoleShadow(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(g, r, tableShadow)
}

func setRoleEel(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(g, r, tableEel)
}

func setRoleNPC(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return setRole(g, r, tableNPC)
}

// Self-Assignable Role
func setSAR(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {

	// Check if the role exists
	_, err = selectSAR(g, r)
	if err == sql.ErrNoRows {

		// Insert if there's none
		return insertSAR(g, r)
	} else if err != nil {
		return nil, err
	}

	return
}
