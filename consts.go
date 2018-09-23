package main

import (
	"os"
)

const (
	rootFolder   = "./Miiko"
	discordPath  = rootFolder + "/discord.json"
	databasePath = rootFolder + "/db.json"
	errorPath    = rootFolder + "/errors.log"
)

const (
	french  = iota
	english = iota
)

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

const (
	permPrivateDirectory os.FileMode = 0700
	permPrivateFile      os.FileMode = 0600
)
