package main

import (
	"database/sql"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// Connection String

func toConnectionString(database Database) string {
	return database.User + ":" + database.Password + "@tcp(" + database.Address + ":" + strconv.Itoa(database.Port) + ")/" + database.Database
}

// Select

func selectRole(g *discordgo.Guild, table string) (id string, err error) {
	err = db.QueryRow("select `role` from ? where server = ?;", table, g.ID).Scan(&id)
	return id, err
}

func selectRoleAdmin(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, "role-administrator")
}

func selectRoleMod(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, "role-moderator")
}

func selectRoleLight(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, "role-light")
}

func selectRoleAbsynthe(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, "role-absynthe")
}

func selectRoleObsidian(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, "role-obsidian")
}

func selectRoleShadow(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, "role-shadow")
}

func selectRoleEel(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, "role-eel")
}

func selectRoleNPC(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, "role-npc")
}

// Insert

func insertRole(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role, table string) (res sql.Result, err error) {
	res, err = tx.Exec("insert into ?(server, role) values(?, ?);", table, g.ID, r.ID)
	return res, err
}

// Update
