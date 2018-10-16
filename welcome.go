package main

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Ask for the guard.
func askForGuard(s *discordgo.Session, g *discordgo.Guild, u *discordgo.User) {

	// Make sure the channel exists
	channel, err := getWelcomeChannel(s, g)
	if err != nil {
		printDiscordError("Couldn't get the channel structure of a welcome channel.", g, nil, nil, u, err)
		return
	}

	// Typing!
	err = s.ChannelTyping(channel.ID)
	if err != nil {
		printDiscordError("Couldn't tell everyone I'm typing some welcome message.", g, channel, nil, u, err)
	}

	if !u.Bot {

		// Ask newcomer what's their guard
		_, err = s.ChannelMessageSend(channel.ID, getWelcomeMessage(u.ID))
		if err != nil {
			printDiscordError("Couldn't welcome a user.", g, channel, nil, u, err)
		}

	} else if u.ID != me.ID {
		setStatus(s, "insulter "+u.Username)

		// Fear the bot!
		_, err = s.ChannelMessageSend(channel.ID, getWelcomeBotMessage(u.ID))
		if err != nil {
			printDiscordError("Couldn't welcome a bot.", g, channel, nil, u, err)
		}
	}
}

func getWelcomeMessage(username string) string {

	// Welcome!
	welcomeList := [...]string{
		"Bonjour <@" + username + ">!",
		"Bonjour, <@" + username + ">.",
		"Bienvenue, <@" + username + ">!",
		"Bienvenue, <@" + username + ">.",
		"Bienvenue à <@" + username + ">!",
		"Bienvenue à toi, <@" + username + ">.",
		"Bienvenue dans notre serveur, <@" + username + ">!",
		"Bienvenue dans notre serveur, <@" + username + ">.",
		"Salutations, <@" + username + ">.",
		"Ah, <@" + username + ">! Nous t'attendions.",
		"<@" + username + ">, tu es là! Je te souhaite la bienvenue.",
		"<@" + username + ">, tu es là! Je te souhaite la bienvenue sur notre serveur.",
		"<@" + username + ">, tu es là! Nous t'attendions.",
		"Ah, voilà <@" + username + ">. Bienvenue!",
		"Ah, voilà <@" + username + ">. Je te souhaite la bienvenue!",
		"Ah, voilà <@" + username + ">. Je te souhaite la bienvenue sur notre serveur.",
		"Ah, voilà <@" + username + ">. Nous t'attendions.",
		"<@" + username + ">, je te souhaite la bienvenue.",
		"<@" + username + ">! Je te souhaite la bienvenue.",
		"<@" + username + ">, je te souhaite la bienvenue sur notre serveur.",
		"<@" + username + ">, nous t'attendions.",
		"Je te souhaite la bienvenue, <@" + username + ">.",
		"Je te souhaite la bienvenue, <@" + username + ">!",
		"Je te souhaite la bienvenue sur notre serveur, <@" + username + ">.",
		"Nous t'attendions, <@" + username + ">.",
		"J'ai le plaisir de vous présenter le nouveau membre du serveur, <@" + username + ">!",
		"J'ai le plaisir de vous présenter le nouveau membre du quartier général, <@" + username + ">!",
		"Souhaitez tous la bienvenue à <@" + username + ">!",
		"Une bonne main d'applaudissement pour <@" + username + ">!",
	}

	// Random
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return welcomeList[random.Intn(len(welcomeList))] + " " + getAskForGuardMessage()
}

func getAskForGuardMessage() string {

	// What's your guard?
	questionList := [...]string{
		"Dans quelle garde es-tu?",
		"Quelle est ta garde?",
		"De quelle garde fais-tu partie?",
		"Peux-tu me dire tu es dans quelle garde?",
		"Peux-tu me dire quelle est ta garde?",
		"Peux-tu me dire de quelle garde tu fais partie?",
		"Dis-moi, dans quelle garde es-tu?",
		"Dis-moi, quelle est ta garde?",
		"Dis-moi, de quelle garde fais-tu partie?",
		"D'ailleurs, dans quelle garde es-tu?",
		"D'ailleurs, quelle est ta garde?",
		"D'ailleurs, de quelle garde fais-tu partie?",
		"Alors, dans quelle garde es-tu?",
		"Alors, quelle est ta garde?",
		"Alors, de quelle garde fais-tu partie?",
	}

	// Random
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return questionList[random.Intn(len(questionList))]
}

func getWelcomeBotMessage(userID string) string {

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Welcome!
	welcomeBotList := [...]string{

		// Wait, what?
		"Mais... <@" + userID + "> est un bot! Qu'est-ce cette chose fait ici?",
		"Mais quel genre de Faery est <@" + userID + ">?",

		// Nope.
		"Non, <@" + userID + ">. Je ne veux pas te voir ici.",
		"Hé, <@" + userID + ">. On ne veut pas de toi ici.",
		"Arrière, <@" + userID + ">!",

		// Botpocalypse
		"T'es venu prendre mon job, <@" + userID + ">?",

		// Passive roast
		"Ça pue, ici! Oh, c'est juste <@" + userID + ">.",
		"Qui vote pour qu'on kick <@" + userID + ">?",
		"On accueille les déchets, maintenant?",
		"Mais quelle abomination!",
		"Beurk.",

		// Notice me senpai!
		"Tiens, un truc moche.",
		"Tiens, un tas de ferraille.",
		"Oh, ça, c'est pas joli.",

		// Community
		"100 PO à celui qui débranche <@" + userID + ">!",
	}
	return welcomeBotList[rand.Intn(len(welcomeBotList))]
}
