package main

import (
	"fmt"
)

type Map [50][150]byte

func update(game *game) {
	x, y := game.player.getPosition()
	game.gameMap[x][y] = '@'
}

func printGame(game *game) {
	printStatus(game.status)
	// printLine(strings.Repeat(" ", 140))
	// printLine(strings.Repeat(" ", 140))
	// printLine(strings.Repeat(" ", 140))
	// printLine(strings.Repeat(" ", 140))
	// printLine(strings.Repeat(" ", 40) + strings.Repeat("-", 50) + strings.Repeat(" ", 50))
	// printLine(strings.Repeat(" ", 40) + "|" + strings.Repeat(".", 48) + "|")
	// printLine(strings.Repeat(" ", 40) + "|" + strings.Repeat(".", 48) + "|")
	// printLine(strings.Repeat(" ", 40) + "|" + strings.Repeat(".", 48) + "|")
	// printLine(strings.Repeat(" ", 40) + "|" + strings.Repeat(".", 48) + "|")
	// printLine(strings.Repeat(" ", 40) + "|" + strings.Repeat(".", 48) + "|")
	// printLine(strings.Repeat(" ", 40) + "|" + strings.Repeat(".", 48) + "|")
	// printLine(strings.Repeat(" ", 40) + "|" + strings.Repeat(".", 48) + "|")
	// printLine(strings.Repeat(" ", 40) + "|" + strings.Repeat(".", 48) + "|")
	// printLine(strings.Repeat(" ", 40) + "|" + strings.Repeat(".", 48) + "|")
	// printLine(strings.Repeat(" ", 40) + "|" + strings.Repeat(".", 48) + "|")
	// printLine(strings.Repeat(" ", 40) + strings.Repeat("-", 50) + strings.Repeat(" ", 50))
	// printLine(strings.Repeat(" ", 140))
	// printLine(strings.Repeat(" ", 140))
	// printLine(strings.Repeat(" ", 140))
	// printLine(strings.Repeat(" ", 140))
	for _, row := range game.gameMap {
		printLine(string(row[:]))
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