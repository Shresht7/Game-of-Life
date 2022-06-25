package main

import (
	"math/rand"
	"time"
)

//	Number of frames to render per second
const FPS = 60

//	The character to represent a single living entity
const ENTITY = "◆"

const (
	ROWS    = 36
	COLUMNS = 120
)

func main() {
	//	Prepare seed for RNG
	rand.Seed(time.Now().Unix())

	grid := NewGrid(ENTITY, ROWS, COLUMNS, true)

	for {
		grid.evolve()
		grid.render()
		time.Sleep((1000 / FPS) * time.Millisecond)
	}
}
