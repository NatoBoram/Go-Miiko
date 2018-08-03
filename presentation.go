package main

import (
	"fmt"
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

	_, err = session.ChannelMessageSend(c.ID, getIntroductionMessage(channel))
	if err != nil {
		fmt.Println("Couldn't ask for introduction.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
	}
}

func getIntroductionMessage(c *discordgo.Channel) string {

	// Welcome!
	list := [...]string{
		"Tu peux te présenter dans <#" + c.ID + ">.",
		"Si tu as quelques instants, je t'invite à te présenter dans <#" + c.ID + ">!",
		"J'aimerais en savoir plus sur toi, tu peux m'en dire plus dans <#" + c.ID + ">?",
		"Je t'invite à faire un tour dans <#" + c.ID + ">!",
		"Maintenant que tu as obtenu ta garde, rends-toi dans <#" + c.ID + ">!",
	}

	// Random
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return list[rand.Intn(len(list))]
}
