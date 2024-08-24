package main

import (
	"fmt"
)

type Map [40][60]Tile

func printGame(game *Game) {
	printStatus(game.status)
	for _, tiles := range game.gameMap {
		var row string
		for _, tile := range tiles {
			row += string(tile.display) + " "
		}
		printLine(row)
	}
	printStats()
	printLine("")
}

func printStatus(s string) {
	printLine(s)
}

func printStats() {
	fmt.Printf("Level: %v\t  ", 1)
	fmt.Printf("HP: %v(%v)\t  ", 10, 20)
	fmt.Printf("Str: %v\t  ", 10)
	fmt.Printf("Exp: %v\t  ", 15)
	printLine("")
}
