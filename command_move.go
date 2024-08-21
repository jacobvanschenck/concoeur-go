package main

func callbackMoveNorth(game *game) error {
	game.status = "moved north"
	x, _ := game.player.getPosition()
	game.player.position.x = x - 1
	return nil
}

func callbackMoveNorthEast(game *game) error {
	game.status = "moved north east"
	x, y := game.player.getPosition()
	game.player.position.x = x - 1
	game.player.position.y = y + 1
	return nil
}

func callbackMoveEast(game *game) error {
	game.status = "moved east"
	_, y := game.player.getPosition()
	game.player.position.y = y + 1
	return nil
}

func callbackMoveSouthEast(game *game) error {
	game.status = "moved south east"
	x, y := game.player.getPosition()
	game.player.position.x = x + 1
	game.player.position.y = y + 1
	return nil
}

func callbackMoveSouth(game *game) error {
	game.status = "moved south"
	x, _ := game.player.getPosition()
	game.player.position.x = x + 1
	return nil
}

func callbackMoveSouthWest(game *game) error {
	game.status = "moved south west"
	x, y := game.player.getPosition()
	game.player.position.x = x + 1
	game.player.position.y = y - 1
	return nil
}

func callbackMoveWest(game *game) error {
	game.status = "moved west"
	_, y := game.player.getPosition()
	game.player.position.y = y - 1
	return nil
}

func callbackMoveNorthWest(game *game) error {
	game.status = "moved north west"
	x, y := game.player.getPosition()
	game.player.position.x = x - 1
	game.player.position.y = y - 1
	return nil
}
