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

// Welcome Channel

// Select Welcome Channel

func selectWelcomeChannel(g *discordgo.Guild) (id string, err error) {
	err = db.QueryRow("select `channel` from `"+tableWelcome+"` where server = ?;", g.ID).Scan(&id)
	return id, err
}

// Insert Welcome Channel

func insertWelcomeChannel(tx *sql.Tx, g *discordgo.Guild, c *discordgo.Channel) (res sql.Result, err error) {
	res, err = tx.Exec("insert into `"+tableWelcome+"`(`server`, `channel`) values(?, ?);", g.ID, c.ID)
	return res, err
}

// Update Welcome Channel

func updateWelcomeChannel(tx *sql.Tx, g *discordgo.Guild, c *discordgo.Channel) (res sql.Result, err error) {
	res, err = tx.Exec("update `"+tableWelcome+"` set `channel` = ? where `server` = ?;", c.ID, g.ID)
	return res, err
}

// Pins

// Select Pin

// Insert Pin

// Update Pin

// Delete Pin

// Roles

// Select Role

func selectRole(g *discordgo.Guild, table string) (id string, err error) {
	err = db.QueryRow("select `role` from `"+table+"` where server = ?;", g.ID).Scan(&id)
	return id, err
}

func selectRoleAdmin(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, tableAdmin)
}

func selectRoleMod(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, tableMod)
}

func selectRoleLight(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, tableLight)
}

func selectRoleAbsynthe(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, tableAbsynthe)
}

func selectRoleObsidian(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, tableObsidian)
}

func selectRoleShadow(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, tableShadow)
}

func selectRoleEel(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, tableEel)
}

func selectRoleNPC(g *discordgo.Guild) (id string, err error) {
	return selectRole(g, tableNPC)
}

// Insert Role

func insertRole(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role, table string) (res sql.Result, err error) {
	return tx.Exec("insert into "+table+"(server, role) values(?, ?);", g.ID, r.ID)
}

func insertRoleAdmin(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(tx, g, r, tableAdmin)
}

func insertRoleMod(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(tx, g, r, tableMod)
}

func insertRoleLight(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(tx, g, r, tableMod)
}

func insertRoleAbsynthe(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(tx, g, r, tableMod)
}

func insertRoleObsidian(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(tx, g, r, tableMod)
}

func insertRoleShadow(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {

	return insertRole(tx, g, r, tableMod)
}
func insertRoleEel(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(tx, g, r, tableMod)
}

func insertRoleNPC(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(tx, g, r, tableMod)
}

// Update Role

func updateRole(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role, table string) (res sql.Result, err error) {
	return tx.Exec("update `"+table+"` set role = ? where server = ;", table, r.ID, g.ID)
}

func updateRoleAdmin(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(tx, g, r, tableAdmin)
}

func updateRoleMod(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(tx, g, r, tableMod)
}

func updateRoleLight(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(tx, g, r, tableLight)
}

func updateRoleAbsynthe(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(tx, g, r, tableAbsynthe)
}

func updateRoleObsidian(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(tx, g, r, tableObsidian)
}

func updateRoleShadow(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(tx, g, r, tableShadow)
}

func updateRoleEel(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(tx, g, r, tableEel)
}

func updateRoleNPC(tx *sql.Tx, g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(tx, g, r, tableNPC)
}
