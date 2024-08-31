package main

import (
	"github.com/jacobvanschenck/concoeur/ecs"
)

type Tile struct {
	position Position
	display  byte
	isSolid  bool
}

type Map struct {
	tiles [40][60]Tile
}

const MapTypeId = "mapTypeId"

func (gameMap Map) typeId() ecs.TypeId {
	return MapTypeId
}

func (gameMap *Map) generateMap() {
	rowIndex := 0

	for _, row := range gameMap.tiles[:13] {
		for j := range row {
			gameMap.tiles[rowIndex][j] = Tile{display: ' '}
		}
		rowIndex++
	}

	for j := range gameMap.tiles[rowIndex] {
		char := &gameMap.tiles[rowIndex][j]
		if j <= 20 || j >= 40 {
			char.display = ' '
		} else {
			char.display = '#'
			char.isSolid = true
		}
	}
	rowIndex++

	for _, row := range gameMap.tiles[rowIndex : rowIndex+10] {
		for j := range row {
			char := &gameMap.tiles[rowIndex][j]
			if j <= 20 || j >= 40 {
				char.display = ' '
			} else if j == 21 || j == 39 {
				char.display = '#'
				char.isSolid = true
			} else {
				char.display = '.'
			}

		}
		rowIndex++
	}

	for j := range gameMap.tiles[rowIndex] {
		char := &gameMap.tiles[rowIndex][j]
		if j <= 20 || j >= 40 {
			char.display = ' '
		} else {
			char.display = '#'
			char.isSolid = true
		}
	}
	rowIndex++

	for i := range gameMap.tiles[rowIndex:] {
		for j := range gameMap.tiles[i] {
			gameMap.tiles[rowIndex][j] = Tile{display: ' '}
		}
		rowIndex++
	}
}
