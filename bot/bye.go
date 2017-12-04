package bot

import (
	"fmt"
	"math/rand"
	"time"

	"../config"
	"github.com/bwmarrin/discordgo"
)

func waitComeBack(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

	// Get guild
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Couldn't get " + m.User.Username + "'s guild ID.")
		fmt.Println(err.Error())
		return
	}

	// Get channel
	channel := config.GetWelcomeChannelByGuildID(guild.ID)
	if channel == "" {
		return
	}

	// Create an invite structure
	var invStruct discordgo.Invite
	invStruct.Temporary = true

	// Create an invite to WelcomeChannel
	var invite *discordgo.Invite
	invite, err = s.ChannelInviteCreate(channel, invStruct)
	if err != nil {
		fmt.Println("Couldn't create an invite in " + guild.Name + ".")
		fmt.Println(err.Error())
		return
	}

	// Bot?
	if m.User.Bot {

		// Typing!
		err = s.ChannelTyping(channel)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Send message
		_, err = s.ChannelMessageSend(channel, getByeBotMessage(m.User.ID))
		if err != nil {
			fmt.Println("Couldn't say bye to " + m.User.Username + "!")
			fmt.Println(err.Error())
		}

	} else {

		// Open channel
		privateChannel, err := s.UserChannelCreate(m.User.ID)
		if err != nil {
			fmt.Println("Couldn't create a private channel with " + m.User.Username + ".")
			fmt.Println(err.Error())
			return
		}

		// Typing!
		err = s.ChannelTyping(privateChannel.ID)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Send message
		_, err = s.ChannelMessageSend(privateChannel.ID, getPrivateByeMessage(invite.Code))
		if err != nil {
			fmt.Println("Couldn't say bye to " + m.User.Username + "!")
			fmt.Println(err.Error())
		}

		// Typing!
		err = s.ChannelTyping(channel)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Announce departure
		_, err = s.ChannelMessageSend(channel, getPublicByeMessage(m.User.ID))
		if err != nil {
			fmt.Println("Couldn't announce the departure of " + m.User.Username + ".")
			fmt.Println(err.Error())
		}
	}
}

func getPrivateByeMessage(inviteCode string) string {

	// Bye Messages
	var byeList []string

	// Messages
	byeList = append(byeList, "Oh, je suis triste de te voir partir! Si tu veux nous rejoindre à nouveau, j'ai créé une invitation pour toi : https://discord.gg/"+inviteCode)
	byeList = append(byeList, "Au revoir! Voici une invitation si tu changes d'idée : https://discord.gg/"+inviteCode)
	byeList = append(byeList, "Tu vas me manquer. Si tu veux me revoir, j'ai créé une invitation pour toi : https://discord.gg/"+inviteCode)

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeList[rand.Intn(len(byeList))]
}

func getPublicByeMessage(userID string) string {

	// Bye Messages
	var byeList []string

	// Messages
	byeList = append(byeList, "J'ai le regret d'annoncer le départ de <@"+userID+">.")
	byeList = append(byeList, "C'est avec émotion que j'annonce le départ de <@"+userID+">.")
	byeList = append(byeList, "L'Oracle a emporté <@"+userID+"> avec elle.")
	byeList = append(byeList, "<@"+userID+"> a quitté la garde.")
	byeList = append(byeList, "Attends, <@"+userID+">, reviens!")
	byeList = append(byeList, "<@"+userID+"> s'est envolé!")
	byeList = append(byeList, "<@"+userID+"> vole de ses propres ailes.")
	byeList = append(byeList, "<@"+userID+"> part à l'aventure!")
	byeList = append(byeList, "L'aventure de <@"+userID+"> se termine ici.")
	byeList = append(byeList, "La garde se souviendra de <@"+userID+">.")
	byeList = append(byeList, "Il pleut lorsque je regarde vers <@"+userID+">.")
	byeList = append(byeList, "Mon coeur se serre à l'annonce du départ de <@"+userID+">.")
	byeList = append(byeList, "<@"+userID+"> a donné sa démission.")

	// Community
	byeList = append(byeList, "Aurevoir, <@"+userID+">. Reviens-nous vite!")
	byeList = append(byeList, "<@"+userID+"> nous a quitté. Souhaiton-lui le meilleur!")
	byeList = append(byeList, "<@"+userID+"> nous a quitté. Elle va nous manquer.")
	byeList = append(byeList, "Adieu, <@"+userID+">! Vole vers d'autres cieux!")
	byeList = append(byeList, "<@"+userID+"> a été transféré vers un autre QG.")
	byeList = append(byeList, "Nous n'oublierons pas le sacrifice de <@"+userID+">!")
	byeList = append(byeList, "Nous avons perdu <@"+userID+">, mais nous restons forts.")

	// Legendary
	byeList = append(byeList, "C'est en ce jour funeste que <@"+userID+"> nous a quitté. Puisse son âme rejoindre le cristal et son héritage mon porte-maanas.")
	// byeList = append(byeList, "<@"+userID+">, en tant qu'ancien membre de notre serveur, a droit a une cérémonie d'adieu. 1. Tu ne dois jamais révéler des informations sensibles sur notre serveur à d'autres aussi longtemps que tu vivras. 2. Tu ne dois jamais utiliser d'anciens contacts rencontrés à travers ta présence dans le serveur pour un gain personnel. 3. Bien que nos chemins puissent avoir divergé, tu dois continuer à vivre ta vie de toutes tes forces, tu ne dois jamais considérer ta propre vie comme quelque chose d'insignifiant, et tu ne dois jamais oublier tes amis aussi longtemps que tu vis.")

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeList[rand.Intn(len(byeList))]
}

func getByeBotMessage(userID string) string {

	// Bye Messages
	var byeBotList []string

	// Messages
	byeBotList = append(byeBotList, "Bon débarras, <@"+userID+">.")
	byeBotList = append(byeBotList, "Bien! Personne ne va s'ennuyer de <@"+userID+">.")
	byeBotList = append(byeBotList, "De toute façon, <@"+userID+"> n'avait aucun lien avec Eldarya.")
	byeBotList = append(byeBotList, "<@"+userID+"> ne nous manquera pas.")
	byeBotList = append(byeBotList, "Ha! <@"+userID+"> est parti. Ça fait plus de popcorn pour moi!")

	// Community
	byeBotList = append(byeBotList, "Nous sommes enfin débarrassés de <@"+userID+">!")
	byeBotList = append(byeBotList, "Oh, <@"+userID+"> est mort. Mais quel dommage.")
	byeBotList = append(byeBotList, "Super! <@"+userID+"> a fiché le camp!")
	byeBotList = append(byeBotList, "Ah? <@"+userID+"> était là?")

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeBotList[rand.Intn(len(byeBotList))]
}
