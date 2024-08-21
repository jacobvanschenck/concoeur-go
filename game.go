package main

import (
	"fmt"
	"os"
)

func startGame(game *game) {
	for game.running {
		clearScreen()
		game.initMap()
		update(game)
		printGame(game)

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
		'y': {
			name:        "y",
			description: "Move north west",
			callback:    callbackMoveNorthWest,
		},
		'k': {
			name:        "k",
			description: "Move north",
			callback:    callbackMoveNorth,
		},
		'u': {
			name:        "u",
			description: "Move north east",
			callback:    callbackMoveNorthEast,
		},
		'l': {
			name:        "l",
			description: "Move east",
			callback:    callbackMoveEast,
		},
		'n': {
			name:        "n",
			description: "Move south east",
			callback:    callbackMoveSouthEast,
		},
		'j': {
			name:        "j",
			description: "Move south",
			callback:    callbackMoveSouth,
		},
		'b': {
			name:        "b",
			description: "Move south west",
			callback:    callbackMoveSouthWest,
		},
		'h': {
			name:        "h",
			description: "Move west",
			callback:    callbackMoveWest,
		},
	}
}

//  y k u
//   \|/
//  h-.-l
//   /|\
//  b j n
