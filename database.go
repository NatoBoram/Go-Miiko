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
	return
}

// Insert Welcome Channel

func insertWelcomeChannel(g *discordgo.Guild, c *discordgo.Channel) (res sql.Result, err error) {
	return db.Exec("insert into `"+tableWelcome+"`(`server`, `channel`) values(?, ?);", g.ID, c.ID)
}

// Update Welcome Channel

func updateWelcomeChannel(g *discordgo.Guild, c *discordgo.Channel) (res sql.Result, err error) {
	return db.Exec("update `"+tableWelcome+"` set `channel` = ? where `server` = ?;", c.ID, g.ID)
}

// Pins

// Select Pin
func selectPin(m *discordgo.Message) (id string, err error) {
	err = db.QueryRow("select `server`, `member`, `message` from `pins` where `message` = ?;", m.ID).Scan(&id)
	return
}

// Insert Pin
func insertPin(g *discordgo.Guild, u *discordgo.User, m *discordgo.Message) (res sql.Result, err error) {
	return db.Exec("insert into `pins`(`server`, `member`, `message`) values(?, ?, ?)", g.ID, u.ID, m.ID)
}

// Update Pin
// Not needed.

// Delete Pin
func deletePin(m *discordgo.Message) (res sql.Result, err error) {
	return db.Exec("delete from `pins` where message = ?;", m.ID)
}

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

func insertRole(g *discordgo.Guild, r *discordgo.Role, table string) (res sql.Result, err error) {
	return db.Exec("insert into "+table+"(server, role) values(?, ?);", g.ID, r.ID)
}

func insertRoleAdmin(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(g, r, tableAdmin)
}

func insertRoleMod(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(g, r, tableMod)
}

func insertRoleLight(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(g, r, tableMod)
}

func insertRoleAbsynthe(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(g, r, tableMod)
}

func insertRoleObsidian(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(g, r, tableMod)
}

func insertRoleShadow(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {

	return insertRole(g, r, tableMod)
}
func insertRoleEel(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(g, r, tableMod)
}

func insertRoleNPC(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return insertRole(g, r, tableMod)
}

// Update Role

func updateRole(g *discordgo.Guild, r *discordgo.Role, table string) (res sql.Result, err error) {
	return db.Exec("update `"+table+"` set role = ? where server = ;", table, r.ID, g.ID)
}

func updateRoleAdmin(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(g, r, tableAdmin)
}

func updateRoleMod(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(g, r, tableMod)
}

func updateRoleLight(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(g, r, tableLight)
}

func updateRoleAbsynthe(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(g, r, tableAbsynthe)
}

func updateRoleObsidian(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(g, r, tableObsidian)
}

func updateRoleShadow(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(g, r, tableShadow)
}

func updateRoleEel(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(g, r, tableEel)
}

func updateRoleNPC(g *discordgo.Guild, r *discordgo.Role) (res sql.Result, err error) {
	return updateRole(g, r, tableNPC)
}

// Presentation Channel

// Select Presentation Channel
func selectPresentationChannel(g *discordgo.Guild) (id string, err error) {
	err = db.QueryRow("select `channel` from `"+tablePresentation+"` where server = ?;", g.ID).Scan(&id)
	return id, err
}

// Insert Presentation Channel
func insertPresentationChannel(g *discordgo.Guild, c *discordgo.Channel) (res sql.Result, err error) {
	return db.Exec("insert into `"+tablePresentation+"`(`server`, `channel`) values(?, ?);", g.ID, c.ID)
}

// Update Presentation Channel
func updatePresentationChannel(g *discordgo.Guild, c *discordgo.Channel) (res sql.Result, err error) {
	return db.Exec("update `"+tablePresentation+"` set `channel` = ? where `server` = ?;", c.ID, g.ID)
}
