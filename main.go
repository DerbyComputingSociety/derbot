package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	//"github.com/bwmarrin/discordgo"
)

var (
	Token string
	Owner string
)

func main() {
	file, err := os.Open("config.txt")
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

	fmt.Printf("%s, %s\n", Owner, Token)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
