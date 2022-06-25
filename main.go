package main

import (
	"math/rand"
	"time"
)

//	Number of frames to render per second
const FPS = 60

//	The character to represent a single living entity
const ENTITY = "◆"

func main() {
	//	Prepare seed for RNG
	rand.Seed(time.Now().Unix())

	grid := NewGrid(80, 120, true)

	for {
		grid.evolve()
		grid.render(ENTITY)
		time.Sleep((1000 / FPS) * time.Millisecond)
	}
}
