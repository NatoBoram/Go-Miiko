package main

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func askForIntroduction(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) (err error) {

	// Get the presentation channel
	channel, err := getPresentationChannel(s, g)
	if err != nil {
		return err
	}

	// Send the introduction message
	_, err = s.ChannelMessageSend(c.ID, getIntroductionMessage(channel))
	return
}

func getIntroductionMessage(c *discordgo.Channel) string {

	// "<#" + c.ID + ">",

	// Welcome!
	list := [...]string{

		// Tu peux
		"Tu peux te présenter dans <#" + c.ID + ">.",
		"Tu peux aller te présenter dans <#" + c.ID + ">.",
		"Tu peux faire un tour dans <#" + c.ID + ">.",
		"Tu peux aller faire un tour dans <#" + c.ID + ">.",
		"Tu peux nous en apprendre plus sur toi dans <#" + c.ID + ">.",
		"Tu peux nous en dire plus sur toi dans <#" + c.ID + ">.",

		// N'hésite pas
		"N'hésite pas à te présenter dans <#" + c.ID + ">.",
		"N'hésite pas à aller te présenter dans <#" + c.ID + ">.",
		"N'hésite pas à faire un tour dans <#" + c.ID + ">.",
		"N'hésite pas à aller faire un tour dans <#" + c.ID + ">.",
		"N'hésite pas à nous en apprendre plus sur toi dans <#" + c.ID + ">.",
		"N'hésite pas à nous en dire plus sur toi dans <#" + c.ID + ">.",

		// Si tu souhaites te présenter, tu peux
		"Si tu souhaites te présenter, tu peux le faire dans <#" + c.ID + ">.",
		"Si tu souhaites te présenter, tu peux faire un tour dans <#" + c.ID + ">.",
		"Si tu souhaites te présenter, tu peux te rendre dans <#" + c.ID + ">.",
		"Si tu souhaites te présenter, passe faire un tour dans <#" + c.ID + ">!",

		// Si tu souhaites, n'hésite pas
		"Si tu souhaites te présenter, n'hésites pas à le faire dans <#" + c.ID + ">.",
		"Si tu souhaites te présenter, n'hésites pas à faire un tour dans <#" + c.ID + ">.",
		"Si tu souhaites te présenter, n'hésites pas à te rendre dans <#" + c.ID + ">.",

		// Si tu souhaites
		"Si tu souhaites te présenter, nous avons un salon de <#" + c.ID + ">.",
		"Si tu souhaites nous en apprendre plus sur toi, tu peux le faire <#" + c.ID + ">.",

		// Je t'invite
		"Je t'invite à faire un tour dans <#" + c.ID + ">.",
		"Si tu as quelques instants, je t'invite à te présenter dans <#" + c.ID + ">.",

		// Reste
		"J'aimerais en savoir plus sur toi, tu peux m'en dire plus dans <#" + c.ID + ">?",
		"Maintenant que tu as obtenu ta garde, tu peux te rendre dans <#" + c.ID + ">!",
		"Le salon <#" + c.ID + "> est là si tu souhaites te présenter.",
		"Le salon <#" + c.ID + "> est à ta disposition si tu souhaites te présenter.",
		"Passe faire un tour dans <#" + c.ID + "> afin que nous en sachions plus sur toi!",
	}

	// Random
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return list[rand.Intn(len(list))]
}
