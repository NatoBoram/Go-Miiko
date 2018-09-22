package main

import (
	"github.com/bwmarrin/discordgo"
)

// Checks if a member has authority over Miiko. Used for administrative purposes.
func canAdministrate(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (has bool, err error) {
	owner := g.OwnerID == m.User.ID
	admin, err := isAdmin(s, g, m)
	return owner || admin, err
}

// Checks if a member has authority over the members. Used for administrative purposes.
func canModerate(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (has bool, err error) {
	owner := g.OwnerID == m.User.ID
	admin, err := isAdmin(s, g, m)
	mod, err := isMod(s, g, m)
	return owner || admin || mod, err
}

// Checks if a member has authority over the members. Used for social purposes.
func isLightOrOver(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (has bool, err error) {
	owner := g.OwnerID == m.User.ID
	admin, err := isAdmin(s, g, m)
	mod, err := isMod(s, g, m)
	light, err := isLight(s, g, m)
	return owner || admin || mod || light, err
}

// isGuardianOrOver checks if a member has a guard. Used for administrative purposes.
func isGuardianOrOver(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (has bool, err error) {
	owner := g.OwnerID == m.User.ID
	admin, err := isAdmin(s, g, m)
	mod, err := isMod(s, g, m)
	light, err := isLight(s, g, m)
	absynthe, err := isAbsynthe(s, g, m)
	obsidian, err := isObsidian(s, g, m)
	shadow, err := isShadow(s, g, m)
	return owner || admin || mod || light || absynthe || obsidian || shadow, err
}

// hasGuard checks if a member has a guard. Used to give people a guard.
func hasGuard(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (has bool, err error) {
	absynthe, err := isAbsynthe(s, g, m)
	obsidian, err := isObsidian(s, g, m)
	shadow, err := isShadow(s, g, m)
	eel, err := isEel(s, g, m)
	npc, err := isNPC(s, g, m)
	return absynthe || obsidian || shadow || eel || npc, err
}
