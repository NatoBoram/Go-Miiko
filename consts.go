package main

import (
	"os"
)

// Paths
const (
	rootFolder   = "./Miiko"
	discordPath  = rootFolder + "/discord.json"
	databasePath = rootFolder + "/db.json"
	errorPath    = rootFolder + "/errors.log"
)

// Languages
const (
	french  = iota
	english = iota
)

// Tables
const (
	tableAdmin    = "role-administrator"
	tableMod      = "role-moderator"
	tableLight    = "role-light"
	tableAbsynthe = "role-absynthe"
	tableObsidian = "role-obsidian"
	tableShadow   = "role-shadow"
	tableEel      = "role-eel"
	tableNPC      = "role-npc"
	tableSAR      = "roles-sar"

	tableWelcome          = "channel-welcome"
	tablePresentation     = "channel-presentation"
	tableMinimumReactions = "minimum-reactions"
)

// Permissions
const (
	permPrivateDirectory os.FileMode = 0700
	permPrivateFile      os.FileMode = 0600
)

// Colors
const (
	colourAdministrator = 0xfee69e
	colourModerator     = 0x96eff5
	colourLight         = colourAdministrator
	colourAbsynthe      = 0x7efb7b
	colourObsidian      = 0xfbb6a8
	colourShadow        = 0xf9b3fc
	colourEel           = 0xbec9f8
	colourNPC           = 0x9c9c9c
	colourBot           = 0x7289da
)
