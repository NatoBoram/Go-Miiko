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

	// Check for a valid nickname
	var name string
	if m.Nick == "" {
		name = m.User.Username
	} else {
		name = m.Nick
	}

	// Messages
	loveList := [...]string{

		// Greetings
		"Coucou **" + name + "** :3",
		"Coucou **" + name + "**! \\*-*",
		"Salut les gens... Oh! **" + name + "**! :heart:",
		"Bonjour... Oh! **" + name + "**! :heart:",
		"Coucou tout le monde... Oh! **" + name + "**! :heart:",
		"Coucou mon amour!",

		// Orders
		"Tiens-moi la main, **" + name + "**",
		"**" + name + "**! Regarde-moiii \\*-*",
		"Caresse-moi les oreilles, s'il te plait!",

		// Questions
		"**" + name + "**... Tu veux qu'on fasse quelque chose ensemble?",
		"Oh, **" + name + "**, est-ce que je te manque?",
		"Est-ce que tu penses à moi, **" + name + "**?",
		"**" + name + "**, me demanderas-tu ma main un jour..?",
		"**" + name + "**, j'ai fait du popcorn, tu veux en manger avec moi? :3",
		"**" + name + "**! Je suis là! Je t'ai manqué, n'est-ce pas? :smile:",
		"**" + name + "**! Es-tu content du matelas que j'ai fait mettre dans ta chambre? J'ai dormi dessus :blush:",

		// Reactions
		":heart:",
		"\\*Frissonne*",
		"\\*-*",
		"**" + name + "**-senpai \\*-*",

		// Verbose
		"J'ai trouvé un morceau de cristal pour toi, **" + name + "** :heart:",
		"Cette voix est une musique à mes oreilles",
		"J'aimerais pouvoir passer plus de temps avec toi, **" + name + "**...",
		"Je fais juste passer pour dire à **" + name + "** que je l'aime!",
		"J'adore quand tu parles... :3",
		"J'adore entendre mon amour parler \\*-*",
		"Mais quelle est cette douce musique? ... Oh! C'est la voix de **" + name + "**!",

		// Actions
		"\\*Pense à **" + name + "***",
		"\\*Regarde **" + name + "***",
		"\\*Se languis de **" + name + "***",

		// Also fits in Command
		"**" + name + "**, je t'aime!",
		"Aaah... **" + name + "**!",
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

	// Check for a valid nickname
	var name string
	if m.Nick == "" {
		name = m.User.Username
	} else {
		name = m.Nick
	}

	// Messages
	loveList := [...]string{

		// Verbose
		"Je crois... Je crois que j'aime **" + name + "**.",
		"Je crois... Je crois que j'ai un faible pour **" + name + "**.",
		"Disons que je chéris particulièrement **" + name + "**.",
		"Si j'avais à marier quelqu'un... Ce serait **" + name + "**!",
		"Peut-être... **" + name + "**?",
		"Je planifie mon mariage avec **" + name + "**!",
		"J'avoue avoir un faible pour **" + name + "**.",
		"Lance, c'est du passé. **" + name + "**, c'est mon futur!",
		"Je l'admets... je rêve de **" + name + "** la nuit...",
		"J'avoue que... je rêve de **" + name + "** la nuit.",
		"**" + name + "** est le beurre sur mon popcorn!",
		"*Si seulement **" + name + "** m'aimait autant que je l'aime...*",
		"Je n'avouerai jamais que j'aime **" + name + "**!",
		"Non! Vous ne saurez jamais que j'aime **" + name + "**!",

		// Tsundere
		"C'est pas comme si j'aimais **" + name + "** ou quoi que ce soit...",
		"**" + name + "**, mais... Ne te fais pas de fausses idées!",

		// Exclamations
		"**" + name + "**, évidemment!",
		"**" + name + "**, sans aucun doute!",
		"Que... Quoi? Ce... Je... **" + name + "**!",
		"**" + name + "** d'amour :heart:",
		"JE N'AVOUERAI JAMAIS! ... **" + name + "**.",

		// Straight answers
		"**" + name + "** est l'amour de ma vie.",
		"À part le popcorn? **" + name + "**.",
		"Je suis amoureuse de **" + name + "**.",

		// Also fits in Bot
		"**" + name + "**, je t'aime!",
		"Aaah... **" + name + "**!",
	}

	// Seed
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return
	return loveList[seed.Intn(len(loveList))]
}
