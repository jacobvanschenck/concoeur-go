package main

func exitAction(game *game, actor *actor) error {
	game.running = false
	return nil
}

func moveAction(game *game, actor *actor, dir direction) error {
	actor.position.x += dir.x
	actor.position.y += dir.y
	return nil
}

func invalidAction(game *game, actor *actor) error {
	game.status = "Invalid command"
	return nil
}
