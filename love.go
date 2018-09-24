package main

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/NatoBoram/Go-Miiko/wheel"
)

func love(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message) bool {

	// Lover
	lover, err := getLover(s, g)
	if err != nil {
		printDiscordError("Couldn't get my lover.", g, c, m, nil, err)
		return false
	}

	// Verify if it's the one true love
	if m.Author.ID == lover.User.ID {

		// Rate Limit
		if wheel.RandomOverPhiPower(100) {
			s.ChannelTyping(c.ID)

			// Give some love!
			_, err = s.ChannelMessageSend(c.ID, getLoveMessage(lover))
			if err != nil {
				printDiscordError("Couldn't express my love.", g, c, m, nil, err)
				return false
			}

			return true
		}
	}
	return false
}

func getLoveMessage(m *discordgo.Member) string {

	// Messages
	loveList := [...]string{

		// Greetings
		"Coucou **" + m.Nick + "** :3",
		"Coucou **" + m.Nick + "**! \\*-*",
		"Salut les gens... Oh! **" + m.Nick + "**! :heart:",
		"Bonjour... Oh! **" + m.Nick + "**! :heart:",
		"Coucou tout le monde... Oh! **" + m.Nick + "**! :heart:",
		"Coucou mon amour!",

		// Orders
		"Tiens-moi la main, **" + m.Nick + "**",
		"**" + m.Nick + "**! Regarde-moiii \\*-*",
		"Caresse-moi les oreilles, s'il te plait!",

		// Questions
		"**" + m.Nick + "**... Tu veux qu'on fasse quelque chose ensemble?",
		"Oh, **" + m.Nick + "**, est-ce que je te manque?",
		"Est-ce que tu penses à moi, **" + m.Nick + "**?",
		"**" + m.Nick + "**, me demanderas-tu ma main un jour..?",
		"**" + m.Nick + "**, j'ai fait du popcorn, tu veux en manger avec moi? :3",
		"**" + m.Nick + "**! Je suis là! Je t'ai manqué, n'est-ce pas? :smile:",
		"**" + m.Nick + "**! Es-tu content du matelas que j'ai fait mettre dans ta chambre? J'ai dormi dessus :blush:",

		// Reactions
		":heart:",
		"\\*Frissonne*",
		"\\*-*",
		"**" + m.Nick + "**-senpai \\*-*",

		// Verbose
		"J'ai trouvé un morceau de cristal pour toi, **" + m.Nick + "** :heart:",
		"Cette voix est une musique à mes oreilles",
		"J'aimerais pouvoir passer plus de temps avec toi, **" + m.Nick + "**...",
		"Je fais juste passer pour dire à **" + m.Nick + "** que je l'aime!",
		"J'adore quand tu parles... :3",
		"J'adore entendre mon amour parler \\*-*",
		"Mais quelle est cette douce musique? ... Oh! C'est la voix de **" + m.Nick + "**!",

		// Actions
		"\\*Pense à **" + m.Nick + "***",
		"\\*Regarde **" + m.Nick + "***",
		"\\*Se languis de **" + m.Nick + "***",

		// Also fits in Command
		"**" + m.Nick + "**, je t'aime!",
		"Aaah... **" + m.Nick + "**!",
	}

	// Seed
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return loveList[seed.Intn(len(loveList))]
}

// GetLoverCmd outputs the lover
func getLoverCmd(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, u *discordgo.User) {
	s.ChannelTyping(c.ID)

	// Get lover
	lover, err := getLover(s, g)
	if err != nil {
		printDiscordError("Couldn't get my lover!", g, c, nil, u, err)
		return
	}

	// Send response
	_, err = s.ChannelMessageSend(c.ID, getLoverMessage(lover))
	if err != nil {
		printDiscordError("Couldn't reveal my lover.", g, c, nil, u, err)
	}
}

func getLoverMessage(m *discordgo.Member) string {

	// Messages
	loveList := [...]string{

		// Verbose
		"Je crois... Je crois que j'aime **" + m.Nick + "**.",
		"Je crois... Je crois que j'ai un faible pour **" + m.Nick + "**.",
		"Disons que je chéris particulièrement **" + m.Nick + "**.",
		"Si j'avais à marier quelqu'un... Ce serait **" + m.Nick + "**!",
		"Peut-être... **" + m.Nick + "**?",
		"Je planifie mon mariage avec **" + m.Nick + "**!",
		"J'avoue avoir un faible pour **" + m.Nick + "**.",
		"Lance, c'est du passé. **" + m.Nick + "**, c'est mon futur!",
		"Je l'admets... je rêve de **" + m.Nick + "** la nuit...",
		"J'avoue que... je rêve de **" + m.Nick + "** la nuit.",
		"**" + m.Nick + "** est le beurre sur mon popcorn!",
		"*Si seulement **" + m.Nick + "** m'aimait autant que je l'aime...*",
		"Je n'avouerai jamais que j'aime **" + m.Nick + "**!",
		"Non! Vous ne saurez jamais que j'aime **" + m.Nick + "**!",

		// Tsundere
		"C'est pas comme si j'aimais **" + m.Nick + "** ou quoi que ce soit...",
		"**" + m.Nick + "**, mais... Ne te fais pas de fausses idées!",

		// Exclamations
		"**" + m.Nick + "**, évidemment!",
		"**" + m.Nick + "**, sans aucun doute!",
		"Que... Quoi? Ce... Je... **" + m.Nick + "**!",
		"**" + m.Nick + "** d'amour :heart:",
		"JE N'AVOUERAI JAMAIS! ... **" + m.Nick + "**.",

		// Straight answers
		"**" + m.Nick + "** est l'amour de ma vie.",
		"À part le popcorn? **" + m.Nick + "**.",
		"Je suis amoureuse de **" + m.Nick + "**.",

		// Also fits in Bot
		"**" + m.Nick + "**, je t'aime!",
		"Aaah... **" + m.Nick + "**!",
	}

	// Seed
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return loveList[seed.Intn(len(loveList))]
}
