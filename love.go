package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"

	"gitlab.com/NatoBoram/Go-Miiko/wheel"
)

func love(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message) bool {

	// Lover
	lover, err := GetLover(db, s, g)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	// Verify if it's the one true love
	if m.Author.ID == lover.ID {

		// Random
		seed := time.Now().UnixNano()
		source := rand.NewSource(seed)
		rand := rand.New(source)
		random := rand.Float64()

		// Rate Limit
		if random < 1/(wheel.Phi()*100) {

			// Member
			member, err := s.GuildMember(g.ID, lover.ID)
			if err != nil {
				fmt.Println("Couldn't get the member I love.")
				fmt.Println("Guild :", g.Name)
				fmt.Println("Channel :", c.Name)
				fmt.Println("User :", m.Author.Username)
				fmt.Println("Message :", m.Content)
				fmt.Println(err.Error())
				return false
			}

			mention := "**" + member.Nick + "**"

			// Give some love!
			s.ChannelTyping(c.ID)
			_, err = s.ChannelMessageSend(c.ID, getLoveMessage(mention))
			if err != nil {
				fmt.Println("Couldn't express my love.")
				fmt.Println("Guild :", g.Name)
				fmt.Println("Channel :", c.Name)
				fmt.Println("User :", m.Author.Username)
				fmt.Println("Message :", m.Content)
				fmt.Println(err.Error())
			}

			return true
		}
	}
	return false
}

func getLoveMessage(name string) string {

	// Messages
	loveList := [...]string{

		// Greetings
		"Coucou " + name + " :3",
		"Coucou " + name + "! \\*-*",
		"Salut les gens... Oh! " + name + "! :heart:",
		"Bonjour... Oh! " + name + "! :heart:",
		"Coucou tout le monde... Oh! " + name + "! :heart:",
		"Coucou mon amour!",

		// Orders
		"Tiens-moi la main, " + name + "",
		"" + name + "! Regarde-moiii \\*-*",
		"Caresse-moi les oreilles, s'il te plait!",

		// Questions
		"" + name + "... Tu veux qu'on fasse quelque chose ensemble?",
		"Oh, " + name + ", est-ce que je te manque?",
		"Est-ce que tu penses à moi, " + name + "?",
		"" + name + ", me demanderas-tu ma main un jour..?",
		"" + name + ", j'ai fait du popcorn, tu veux en manger avec moi? :3",
		"" + name + "! Je suis là! Je t'ai manqué, n'est-ce pas? :smile:",
		"" + name + "! Es-tu content du matelas que j'ai fait mettre dans ta chambre? J'ai dormi dessus :blush:",

		// Reactions
		":heart:",
		"\\*Frissonne*",
		"\\*-*",
		"" + name + "-senpai \\*-*",

		// Verbose
		"J'ai trouvé un morceau de cristal pour toi, " + name + " :heart:",
		"Cette voix est une musique à mes oreilles",
		"J'aimerais pouvoir passer plus de temps avec toi, " + name + "...",
		"Je fais juste passer pour dire à " + name + " que je l'aime!",
		"J'adore quand tu parles... :3",
		"J'adore entendre mon amour parler \\*-*",
		"Mais quelle est cette douce musique? ... Oh! C'est la voix de " + name + "!",

		// Actions
		"\\*Pense à " + name + "*",
		"\\*Regarde " + name + "*",
		"\\*Se languis de " + name + "*",

		// Also fits in Command
		"" + name + ", je t'aime!",
		"Aaah... " + name + "!",
	}

	// Seed
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)

	// Return
	return loveList[seed.Intn(len(loveList))]
}

// GetLoverCmd outputs the lover
func GetLoverCmd(db *sql.DB, s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, u *discordgo.User) {

	// Inform the user that I'm typing
	s.ChannelTyping(c.ID)

	// Get lover
	lover, err := GetLover(db, s, g)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var mention string
	if u.ID == g.OwnerID {
		mention = "<@" + lover.ID + ">"
	} else {

		// Don't mention because we don't want to spam the lover
		member, err := s.GuildMember(g.ID, lover.ID)
		if err != nil {
			fmt.Println("Couldn't get the member I love.")
			fmt.Println(err.Error())
			return
		}
		mention = "**" + member.Nick + "**"
	}

	// Send response
	_, err = s.ChannelMessageSend(c.ID, getLoverMessage(mention))
	if err != nil {
		fmt.Println("Couldn't reveal my lover.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		fmt.Println("User :", u.Username)
	}
}

// GetLover gets this guild's lover.
func GetLover(db *sql.DB, s *discordgo.Session, g *discordgo.Guild) (*discordgo.User, error) {

	var (
		userID string
		pins   int
	)

	// Select potential lovers
	rows, err := db.Query("select `member`, `count` from `pins-count` where `server` = ? order by `count` desc;", g.ID)
	if err != nil {
		fmt.Println("Couldn't get my lovers from this guild.")
		fmt.Println("Guild :", g.Name)
		return nil, err
	}
	defer rows.Close()

	// For each rows
	for rows.Next() {

		// Scan it
		err := rows.Scan(&userID, &pins)
		if err != nil {
			fmt.Println("Couldn't scan a potential lover.")
			fmt.Println("Guild :", g.Name)
			continue
		}

		// Member
		member, err := s.GuildMember(g.ID, userID)
		if err != nil {
			fmt.Println("Couldn't get a potential lover's member.")
			fmt.Println("Guild :", g.Name)
			continue
		}

		// Owner
		if g.OwnerID == member.User.ID {
			continue
		}

		// Roles
		if len(member.Roles) == 1 {
			return member.User, nil
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("Couldn't loop my lovers.")
		fmt.Println("Guild :", g.Name)
		return nil, err
	}

	// Unreachable code.
	user, err := s.User(g.OwnerID)
	if err != nil {
		fmt.Println("Couldn't love the owner.")
		fmt.Println("Guild :", g.Name)
		return nil, err
	}

	return user, err
}

func getLoverMessage(name string) string {

	// Messages
	loveList := [...]string{

		// Verbose
		"Je crois... Je crois que j'aime " + name + ".",
		"Je crois... Je crois que j'ai un faible pour " + name + ".",
		"Disons que je chéris particulièrement " + name + ".",
		"Si j'avais à marier quelqu'un... Ce serait " + name + "!",
		"Peut-être... " + name + "?",
		"Je planifie mon mariage avec " + name + "!",
		"J'avoue avoir un faible pour " + name + ".",
		"Lance, c'est du passé. " + name + ", c'est mon futur!",
		"Je l'admets... je rêve de " + name + " la nuit...",
		"J'avoue que... je rêve de " + name + " la nuit.",
		"" + name + " est le beurre sur mon popcorn!",
		"*Si seulement " + name + " m'aimait autant que je l'aime...*",
		"Je n'avouerai jamais que j'aime " + name + "!",
		"Non! Vous ne saurez jamais que j'aime " + name + "!",

		// Tsundere
		"C'est pas comme si j'aimais " + name + " ou quoi que ce soit...",
		"" + name + ", mais... Ne te fais pas de fausses idées!",

		// Exclamations
		"" + name + ", évidemment!",
		"" + name + ", sans aucun doute!",
		"Que... Quoi? Ce... Je... " + name + "!",
		"" + name + " d'amour :heart:",
		"JE N'AVOUERAI JAMAIS! ... " + name + ".",

		// Straight answers
		"" + name + " est l'amour de ma vie.",
		"À part le popcorn? " + name + ".",
		"Je suis amoureuse de " + name + ".",

		// Also fits in Bot
		"" + name + ", je t'aime!",
		"Aaah... " + name + "!",
	}

	// Seed
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)

	// Return
	return loveList[seed.Intn(len(loveList))]
}
