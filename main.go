package main

import (
	"time"
)

const FPS = 60

func main() {
	grid := NewGrid(80, 120, true)

	for {
		grid.evolve()
		grid.render("O")
		time.Sleep((1000 / FPS) * time.Millisecond)
	}
}
