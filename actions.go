package main

func exitAction(game *Game, actor *Actor) error {
	game.running = false
	return nil
}

func moveAction(game *Game, actor *Actor, dir Direction) error {
	actor.position.x += dir.x
	actor.position.y += dir.y
	return nil
}

func invalidAction(game *Game, actor *Actor) error {
	game.status = "Invalid command"
	return nil
}
