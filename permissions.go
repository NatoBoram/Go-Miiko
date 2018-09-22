package main

import (
	"github.com/bwmarrin/discordgo"
)

// hasAuthority checks if a member has authority over the members. Used for social purpose.
func hasAuthority(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (has bool, err error) {
	owner := g.OwnerID == m.User.ID
	admin, err := isAdmin(s, g, m)
	mod, err := isMod(s, g, m)
	light, err := isLight(s, g, m)
	return owner || admin || mod || light, err
}

// hasGuard checks if a member has a guard.
func hasGuard(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) (has bool, err error) {
	absynthe, err := isAbsynthe(s, g, m)
	obsidian, err := isObsidian(s, g, m)
	shadow, err := isShadow(s, g, m)
	eel, err := isEel(s, g, m)
	npc, err := isNPC(s, g, m)
	return absynthe || obsidian || shadow || eel || npc, err
}
