package main

type Action struct {
	perform     func(game *game, actor *actor) error
	name        string
	description string
}

type position struct {
	x int
	y int
}

type direction struct {
	x int
	y int
}

type actor struct {
	position position
	takeTurn func() Action
	display  byte
}

func getActions() map[byte]Action {
	return map[byte]Action{
		'q': {
			name:        "q",
			description: "Quit the game",
			perform:     exitAction,
		},
		'y': {
			name:        "y",
			description: "Move north west",
			perform: func(game *game, actor *actor) error {
				err := moveAction(game, actor, direction{x: -1, y: -1})
				return err
			},
		},
		'k': {
			name:        "k",
			description: "Move north",
			perform: func(game *game, actor *actor) error {
				err := moveAction(game, actor, direction{x: -1, y: 0})
				return err
			},
		},
		'u': {
			name:        "u",
			description: "Move north east",
			perform: func(game *game, actor *actor) error {
				err := moveAction(game, actor, direction{x: -1, y: 1})
				return err
			},
		},
		'l': {
			name:        "l",
			description: "Move east",
			perform: func(game *game, actor *actor) error {
				err := moveAction(game, actor, direction{x: 0, y: 1})
				return err
			},
		},
		'n': {
			name:        "n",
			description: "Move south east",
			perform: func(game *game, actor *actor) error {
				err := moveAction(game, actor, direction{x: 1, y: 1})
				return err
			},
		},
		'j': {
			name:        "j",
			description: "Move south",
			perform: func(game *game, actor *actor) error {
				err := moveAction(game, actor, direction{x: 1, y: 0})
				return err
			},
		},
		'b': {
			name:        "b",
			description: "Move south west",
			perform: func(game *game, actor *actor) error {
				err := moveAction(game, actor, direction{x: 1, y: -1})
				return err
			},
		},
		'h': {
			name:        "h",
			description: "Move west",
			perform: func(game *game, actor *actor) error {
				err := moveAction(game, actor, direction{x: 0, y: -1})
				return err
			},
		},
	}
}

// type melee struct {
// 	minDamage int
// 	maxDamage int
// }
//
// func (melee melee) hit() {
// 	// do stuff
// }
//
// type defense struct {
// 	armor      int
// 	dodgeBonus int
// }
//
// func (defense defense) defend() {
// 	// do stuff
// }
//
// type item struct {
// 	melee   melee
// 	ranged  melee
// 	defense defense
// }
