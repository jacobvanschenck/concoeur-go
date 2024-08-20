package main

import (
	"os"

	"golang.org/x/term"
)

type game struct {
	// gameMap Map
	running bool
	status  string
}

func main() {
	// Get the file descriptor for stdin
	fd := int(os.Stdin.Fd())

	game := game{
		running: true,
	}

	// Save the current state of the terminal
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState) // Restore the terminal when done

	startGame(&game)
}
