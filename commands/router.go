package commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const _DELIM = " "

type CommandContext struct {
	Command string
	Args    []string
	Session *discordgo.Session
	Message *discordgo.MessageCreate
}

type Command struct {
	Name        string
	Description string
	HandlerFunc func(context CommandContext)
}
type Commands []Command

func Handle(command string, s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf(
		"%s:\t%s\n",
		m.Author.Username,
		command)

	cmdSplit := strings.Split(command, _DELIM)
	cmdName := cmdSplit[0]
	cmdArgs := cmdSplit[1:]

	context := CommandContext{
		cmdName,
		cmdArgs,
		s,
		m,
	}

	for _, command := range commands {
		if command.Name == cmdName {
			command.HandlerFunc(context)
		}
	}
}

var commands = Commands{
	Command{
		"help",
		"display information about the available commands",
		help,
	},
}

func help(context CommandContext) {
	log.Println("show help")
}
