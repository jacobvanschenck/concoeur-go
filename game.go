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

	game.drawMap()
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
		game.gameMap[actor.position.x][actor.position.y] = Tile{display: actor.display}
	}
	return game
}

func (game *Game) drawMap() {
	var rows Map
	rowIndex := 0

	for _, row := range rows[:13] {
		for j := range row {
			rows[rowIndex][j] = Tile{display: ' '}
		}
		rowIndex++
	}

	for j := range rows[rowIndex] {
		char := &rows[rowIndex][j]
		if j <= 20 || j >= 40 {
			char.display = ' '
		} else {
			char.display = '#'
			char.isSolid = true
		}
	}
	rowIndex++

	for _, row := range rows[rowIndex : rowIndex+10] {
		for j := range row {
			char := &rows[rowIndex][j]
			if j <= 20 || j >= 40 {
				char.display = ' '
			} else if j == 21 || j == 39 {
				char.display = '#'
				char.isSolid = true
			} else {
				char.display = '.'
			}

		}
		rowIndex++
	}

	for j := range rows[rowIndex] {
		char := &rows[rowIndex][j]
		if j <= 20 || j >= 40 {
			char.display = ' '
		} else {
			char.display = '#'
			char.isSolid = true
		}
	}
	rowIndex++

	for i := range rows[rowIndex:] {
		for j := range rows[i] {
			rows[rowIndex][j] = Tile{display: ' '}
		}
		rowIndex++
	}

	game.gameMap = rows
}

func startGame(game *Game) {
	clearScreen()
	printGame(game)
	for game.running {
		game.drawMap()
		for i := range game.actors {
			actor := &game.actors[i]
			action := actor.takeTurn()
			action.perform(game, actor)
			game.gameMap[actor.position.x][actor.position.y] = Tile{display: actor.display}
		}
		clearScreen()
		printGame(game)
	}

	printLine("Thanks for playing")
}
