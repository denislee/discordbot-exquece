package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/joho/godotenv/autoload"

	"github.com/bwmarrin/discordgo"
	"github.com/nanobox-io/golang-scribble"
)

// invite link : https://discord.com/oauth2/authorize?client_id=858208414379409449&scope=bot+applications.commands

// bot list:  https://discord.com/developers/applications

type Counter struct {
	Current string
}

var (
	Token   string
	current int
)

func main() {

	botToken := os.Getenv("BOT_TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
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
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "exquece" {
		db, _ := scribble.New("./db", nil)
		counter := Counter{}
		db.Read("counter", "counter", &counter)
		current, _ := strconv.Atoi(counter.Current)
		current++
		db.Write("counter", "counter", Counter{Current: strconv.Itoa(current)})
		s.ChannelMessageSend(m.ChannelID, "exquece: "+strconv.Itoa(current))
	}
}
