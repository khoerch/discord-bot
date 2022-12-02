package main

import (
    // "encoding/json"
    // "flag"
    "fmt"
    // "io/ioutil"
    // "net/http"
    "os"
    "os/signal"
    // "strings"
    "syscall"

    "github.com/bwmarrin/discordgo"
		"github.com/joho/godotenv"
)

func main() {

	// Load details from .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Get bot token from .env file
	Token := os.Getenv("BOT_TOKEN")

	// Create a discordgo session
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord Session", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
			fmt.Println("error opening connection,", err)
			return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!bread" {

		var err error
		_, err = s.ChannelMessageSend(m.ChannelID, ":bread: get that bread :bread:")
		if err != nil {
			fmt.Println(err)
		}

	}

}