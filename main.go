package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// GetToken from either TOKEN or TOKEN_FILE environment variables.
// There are two options for passing the secret token to the program:
// 1. TOKEN contains the secret token directly and should not be used in production, but it's fine for development.
// 2. TOKEN_FILE contains the full path to a file that contains the secret.
// TOKEN_FILE is preferred since it can be more secure.
func GetToken() (token string) {
	if value, set := os.LookupEnv("TOKEN"); set {
		log.Println("Found TOKEN, using that")
		token = value
	} else if file, set := os.LookupEnv("TOKEN_FILE"); set {
		log.Println("Found TOKEN_FILE, reading that")
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal("Failed to read token file")
		}
		token = strings.TrimSpace(string(data))
	} else {
		log.Fatal("No token found. Set TOKEN or TOKEN_FILE")
	}

	return token
}

func main() {
	// Get the bot token and open a discord session with it
	token := GetToken()
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Failed to create a discord session\n", err)
	}

	// Add event handler
	session.AddHandler(messageCreate)

	// Specify minimal gateway intents
	// See https://discord.com/developers/docs/topics/gateway#gateway-intents
	session.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

	// Open the websocket that facilitates communication with discord
	err = session.Open()
	if err != nil {
		log.Fatal("Failed to open a websocket for communicating with discord\n", err)
	}
	defer session.Close()

	// Wait for signal from the os before exiting
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop
}

// messageCreate handles messages being sent in any channel the bot has access to.
func messageCreate(s *discordgo.Session, msg *discordgo.MessageCreate) {
	// Ignore messages sent by the bot itself
	if msg.Author.ID == s.State.User.ID {
		return
	}

	s.ChannelMessageSend(msg.ChannelID, "STFU!")
}
