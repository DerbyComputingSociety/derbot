package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	//"github.com/bwmarrin/discordgo"
)

var (
	Token string
	Owner string
)

func main() {
	const input = "token=123123\nowner=monodokimes#1072"
	reader := strings.NewReader(input)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()
		values := strings.Split(text, "=")
		for i := 0; i < 2; i++ {
			key := values[0]
			value := values[1]

			switch key {
			case "token":
				Token = value
			case "owner":
				Owner = value
			default:
				fmt.Printf("unknown key %s", key)
			}
		}
	}

	fmt.Printf("%s, %s\n", Owner, Token)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
