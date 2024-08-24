package main

func exitAction(game *Game, actor *Actor) error {
	game.running = false
	return nil
}

func moveAction(game *Game, actor *Actor, dir Direction) error {
	game.status = ""
	newPosition := Position{x: actor.position.x + dir.x, y: actor.position.y + dir.y}
	tile := game.gameMap[newPosition.x][newPosition.y]
	if tile.isSolid {
		game.status = "There is a wall there"
		return nil
	}
	actor.position.x += dir.x
	actor.position.y += dir.y
	return nil
}

func invalidAction(game *Game, actor *Actor) error {
	game.status = "Invalid command"
	return nil
}
