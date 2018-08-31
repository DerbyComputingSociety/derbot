package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var _config Config

func main() {
	_config = ReadConfig()

	discord, err := discordgo.New("Bot " + _config.Token)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

// Handle a message created on any channel the bot has access to
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, _config.Prefix) {
		log.Printf(
			"%s\t%s\t%s",
			m.ChannelID,
			m.Author.ID,
			strings.TrimPrefix(m.Content, _config.Prefix))

		s.ChannelMessageSend(m.ChannelID, "command used")
	}
}
