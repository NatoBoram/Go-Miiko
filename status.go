package main

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func statusLoop(s *discordgo.Session) {
	for {
		updateStatus(s)
		time.Sleep(time.Hour)
	}
}

func updateStatus(s *discordgo.Session) {

	// Random
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Declare choices
	verbs := [...]string{
		"administrer",
		"modérer",
		"prendre soin de",
	}
	guilds := s.State.Guilds

	// Pick one
	verb := verbs[random.Intn(len(verbs))]
	guild := guilds[random.Intn(len(guilds))]

	// Hey I'm in this guild!
	s.UpdateStatus(0, verb+" "+guild.Name)
}
