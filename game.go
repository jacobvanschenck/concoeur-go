package main

import (
	"os"
)

// Map is going to be a resourece

type Game struct {
	gameMap Map
	actors  []Actor
	running bool
	status  string
}

func newGame() Game {
	game := Game{}

	game.gameMap.generateMap()
	game.running = true
	game.actors = []Actor{
		{
			position: Position{
				x: 19,
				y: 29,
			},
			display: '@',
			takeTurn: func() Action {
				b := make([]byte, 1)
				os.Stdin.Read(b)

				key := b[0]

				availableActions := getActions()
				action, ok := availableActions[key]

				if !ok {
					return Action{perform: invalidAction}
				}

				return action
			},
		},
	}
	for _, actor := range game.actors {
		game.gameMap.tiles[actor.position.x][actor.position.y] = Tile{display: actor.display}
	}
	return game
}

func startGame(game *Game) {
	clearScreen()
	printGame(game)
	for game.running {
		game.gameMap.generateMap()
		for i := range game.actors {
			actor := &game.actors[i]
			action := actor.takeTurn()
			action.perform(game, actor)
			game.gameMap.tiles[actor.position.x][actor.position.y] = Tile{display: actor.display}
		}
		clearScreen()
		printGame(game)
	}

	printLine("Thanks for playing")
}
