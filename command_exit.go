package main

func callbackExit(game *game) error {
	game.running = false
	return nil
}
