package main

import (
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/NatoBoram/Go-Miiko/wheel"

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

// Add +1 to the minimum reactions of a channel.
func addMinimumReactions(c *discordgo.Channel) (err error) {
	min, err := selectMinimumReactions(c)
	if err == sql.ErrNoRows {
		insertMinimumReactions(c, pinAbsoluteMinimum+1)
	} else if err == nil {
		updateMinimumReactions(c, wheel.MaxInt(pinAbsoluteMinimum, min+1))
	}
	return
}

func setStatus(s *discordgo.Session, status string) (err error) {

	// Insert the status in the database
	res, err := insertStatus(status)
	if err != nil {
		return
	}

	// Apply it!
	refreshStatus(s)

	// Get the last inserted ID
	id, err := res.LastInsertId()
	if err != nil {
		return
	}

	go func() {

		// Wait 60 seconds. Yeah, it's hard-coded.
		time.Sleep(60 * time.Second)

		// Delete the status
		_, err = deleteStatus(int(id))
		if err != nil {
			fmt.Println("Couldn't delete a status from the database.")
			fmt.Println(err.Error())
		}

		// Pick-up another status
		refreshStatus(s)
	}()

	return
}
