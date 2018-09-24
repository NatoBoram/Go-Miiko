package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"gitlab.com/NatoBoram/Go-Miiko/wheel"
)

// Popcorn command
func popcorn(s *discordgo.Session, c *discordgo.Channel, m *discordgo.Message) (done bool) {

	// Check for popcorn keywords
	contains := false
	for _, keyword := range popcornStrings {
		if strings.Contains(strings.ToLower(m.Content), keyword) {
			contains = true
			break
		}
	}

	// Return if there's no popcorn in there
	if !contains {
		return false
	}

	// Increase the power of Phi to reduce chances.
	if wheel.RandomOverPhiPower(1) {

		// It's popcorn time!
		s.ChannelTyping(c.ID)
		_, err := s.ChannelMessageSend(c.ID, getPopcornMessage())
		if err != nil {
			printDiscordError("Couldn't tell everyone how much I love popcorn. Sad :(", nil, c, m, nil, err)
			return false
		}

		// Popcorn was successful!
		return true
	}

	// Random chance decided it wasn't popcorn time.
	return false
}

func getPopcornMessage() string {

	// Popcorn Messages
	popcornList := [...]string{

		// Exclamation
		"Popcorn?",
		"Popcorn!",
		"Popcorn?!",
		"Popcorn!?",
		"**Popcorn?!**",
		"**Popcorn!?**",
		"Ah, popcorn!",
		"Hmm, du popcorn...",
		"Hmm, du *popcorn*...",

		// Question
		"On parle de popcorn?",
		"Quelqu'un a dit popcorn?",
		"Quelqu'un a dit popcorn?!",
		"Quelqu'un a dit **popcorn**?!",
		"Quelqu'un a parlé de popcorn?",
		"Quelqu'un a parlé de popcorn?!",
		"Ai-je bien entendu popcorn?",
		"Ai-je bien entendu popcorn?!",
		"Ai-je bien entendu **popcorn**?!",

		// WTF Miiko
		"Moi, j'aime le popcorn!",
		"Le popcorn, c'est génial!",

		// Uhh...
		"Le popcorn, c'est bon et tout, mais il ne faut pas oublier les friandises. J'adore les friandises!",
		"Imagine si... On mélangeait du popcorn... Avec des friandises!",
	}

	// Seed
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return popcornList[random.Intn(len(popcornList))]
}
