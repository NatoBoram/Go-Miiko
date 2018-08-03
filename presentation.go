package main

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func askForIntroduction(g *discordgo.Guild, c *discordgo.Channel) {

	channel, err := getPresentationChannel(g)
	if err != nil {
		// There's no presentation channel.
		return
	}

	session.ChannelMessageSend(c.ID, getIntroductionMessage(channel))

}

func getIntroductionMessage(c *discordgo.Channel) string {

	// Welcome!
	list := [...]string{
		"Tu peux te présenter dans <#" + c.ID + ">.",
	}

	// Random
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return list[rand.Intn(len(list))]
}
