package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func createStack(info string, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, u *discordgo.User, err error) (stack string) {

	if g != nil {
		info += "\nGuild : " + g.Name
	}

	if c != nil {
		info += "\nChannel : " + c.Name
	}

	if m != nil {
		info += "\nMessage : " + m.Content
		info += "\nAuthor : " + m.Author.Username
	}

	if u != nil {
		info += "\nUser : " + u.Username
	}

	if err != nil {
		info += "\nError : " + err.Error()
	}

	info += "\n"

	return info
}

func logError(log string) {

	// Create required directories
	os.MkdirAll(rootFolder, permPrivateDirectory)

	// Open the file in append mode
	file, err := os.OpenFile(errorPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, permPrivateFile)
	if err != nil {
		fmt.Println("Couldn't open the error log file.")
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	// Append the error log
	if _, err = file.WriteString(log); err != nil {
		fmt.Println("Couldn't write the error log.")
		fmt.Println(err.Error())
		return
	}
}

func printDiscordError(info string, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message, u *discordgo.User, err error) {
	stack := createStack(info, g, c, m, u, err)
	fmt.Println(stack)
	logError(stack)
}
