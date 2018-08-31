package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const _PATH = "config.txt"
const _TEMPLATE = "token=\nprefix\n"
const _SPLITTER = "="

type Config struct {
	Token  string
	Prefix string
}

func ReadConfig() Config {
	var config Config

	file, err := os.Open(_PATH)
	defer file.Close()

	// If there's an error the file doesn't exist
	if err != nil {
		log.Println("Config not found, creating...")
		file, err = os.Create(_PATH)
		check(err)

		_, err := file.WriteString(_TEMPLATE)
		check(err)
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" || !strings.Contains(text, _SPLITTER) {
			continue
		}

		values := strings.Split(text, _SPLITTER)

		key := values[0]
		value := values[1]

		switch key {
		case "token":
			config.Token = value
		case "prefix":
			config.Prefix = value
		default:
			log.Printf("unknown key '%s'\n", key)
		}
	}

	if config.Token == "" || config.Prefix == "" {
		log.Fatalf("Config not set. Please fill in %s\n", _PATH)
	}

	check(scanner.Err())

	return config
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
