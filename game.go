package main

import (
	"os"
)

type Game struct {
	gameMap Map
	actors  []Actor
	running bool
	status  string
}

func newGame() Game {
	game := Game{}

	game.initMap()
	game.running = true
	game.actors = []Actor{
		{
			position: Position{
				x: 25,
				y: 75,
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

				// err := command.callback(game)
				// if err != nil {
				// 	fmt.Println(err)
				// }
			},
		},
	}
	for _, actor := range game.actors {
		game.gameMap[actor.position.x][actor.position.y] = actor.display
	}
	return game
}

func (game *Game) initMap() {
	var rows Map

	for i := range rows {
		for j := range rows[i] {
			rows[i][j] = '.'
		}
	}

	game.gameMap = rows
}

func startGame(game *Game) {
	clearScreen()
	printGame(game)
	for game.running {
		game.initMap()
		for i := range game.actors {
			actor := &game.actors[i]
			action := actor.takeTurn()
			action.perform(game, actor)
			game.gameMap[actor.position.x][actor.position.y] = actor.display
		}
		clearScreen()
		printGame(game)
	}

	printLine("Thanks for playing")
}
