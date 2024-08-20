package main

import (
	"fmt"
	"os"
)

func startGame(game *game) {
	for game.running {
		clearScreen()
		printMap(game)

		b := make([]byte, 1)
		os.Stdin.Read(b)

		key := b[0]

		availableCommands := getCommands()
		command, ok := availableCommands[key]

		if !ok {
			game.status = "Invalid command"
			continue
		}

		err := command.callback(game)
		if err != nil {
			fmt.Println(err)
		}
	}

	clearScreen()
	printLine("Thanks for playing")
}

type command struct {
	name        string
	description string
	callback    func(*game) error
}

func getCommands() map[byte]command {
	return map[byte]command{
		'q': {
			name:        "q",
			description: "Quit the game",
			callback:    callbackExit,
		},
		'g': {
			name:        "g",
			description: "Gold",
			callback:    callbackGold,
		},
	}
}
