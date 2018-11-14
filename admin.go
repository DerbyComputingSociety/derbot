package main

import (
	"log"
)

func Help(context CommandContext) {
	user := context.Message.Author
	dmChannel, err := context.Session.UserChannelCreate(user.ID)
	if err != nil {
		log.Print(err)
		return
	}

	content := "help"

	_, err = context.Session.ChannelMessageSend(dmChannel.ID, content)
	if err != nil {
		log.Print(err)
		return
	}
}
