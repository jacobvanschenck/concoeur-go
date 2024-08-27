package main

type Position struct {
	x int
	y int
}

const PositionTypeId = "positionTypeId"

func (position *Position) typeId() string {
	return PositionTypeId
}

// Player

// Monster

// Health

// Wall

// Display { position, character }
