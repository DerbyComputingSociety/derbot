package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const CONFIG_PATH = "config.txt"

var (
	Token string
	Owner string
)

func main() {
	readConfig()

	discord, err := discordgo.New("Bot " + Token)
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

	prefix := "::"

	if strings.HasPrefix(m.Content, prefix) {
		log.Printf(
			"%s\t%s\t%s",
			m.ChannelID,
			m.Author.ID,
			strings.TrimPrefix(m.Content, prefix))

		s.ChannelMessageSend(m.ChannelID, "command used")
	}
}

func readConfig() {
	file, err := os.Open(CONFIG_PATH)
	defer file.Close()

	if err != nil {
		log.Println("Config not found, creating...")

		if _, err := file.WriteString("token=\nowner=\n"); err != nil {
			log.Fatal(err)
		}

		log.Println("Please fill in the config file and run again")
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()
		values := strings.Split(text, "=")

		key := values[0]
		value := values[1]

		switch key {
		case "token":
			Token = value
		case "owner":
			Owner = value
		default:
			fmt.Printf("unknown key '%s'\n", key)
		}
	}

	if Token == "" || Owner == "" {
		log.Fatalln("Config not set. Please fill in config.txt")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
