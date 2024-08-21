package main

type position struct {
	x int
	y int
}

type game struct {
	gameMap Map
	player  player
	running bool
	status  string
}

func newGame() game {
	game := game{}

	game.initMap()
	game.running = true
	game.player = player{
		position: position{
			x: 25,
			y: 75,
		},
	}
	return game
}

func (game *game) initMap() {
	var rows Map

	for i := range rows {
		for j := range rows[i] {
			rows[i][j] = '.'
		}
	}

	game.gameMap = rows
}

type player struct {
	position position
}

func (player *player) getPosition() (int, int) {
	x := player.position.x
	y := player.position.y
	return x, y
}
