package main

import (
	"fmt"
	"math/rand"
	"os"
)

//	====
//	GRID
//	====

//	The game of life grid
type Grid struct {
	rows    int
	columns int
	matrix  [][]int
}

//	Create a new grid
func NewGrid(rows, columns int, randomize bool) *Grid {
	grid := new(Grid) //	Instantiate a new Grid

	grid.rows = rows
	grid.columns = columns
	grid.matrix = make([][]int, columns) //	Create grid columns
	for c := range grid.matrix {
		grid.matrix[c] = make([]int, rows) //	Create grid rows

		//	Randomize values (0 and 1)
		if randomize {
			for r := range grid.matrix[c] {
				grid.matrix[c][r] = rand.Intn(2)
			}
		}

	}

	return grid
}

//	Count the number of live neighbours on the grid for given x and y positions
func (g *Grid) countNeighbours(xPos, yPos int) int {
	count := 0 //	Tracks the count of live neighbours

	//	Iterate over the immediate neighbours
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			posX := (g.columns + xPos + i) % g.columns //	Track X position on the grid
			posY := (g.rows + yPos + j) % g.rows       //	Track Y position on the grid
			count += g.matrix[posX][posY]
		}
	}

	//	Subtract the center's value (it doesn't count as a neighbour)
	count -= g.matrix[xPos][yPos]

	return count
}

//	Evolve the grid according to the rules of the game of life
func (g *Grid) evolve() {
	nextGen := NewGrid(g.rows, g.columns, false) //	Initialize a grid to store the state of the next generation

	for c := 0; c < g.columns; c++ {
		for r := 0; r < g.rows; r++ {

			liveNeighbours := g.countNeighbours(c, r) //	Get neighbour count
			state := g.matrix[c][r]                   //	Track the current state of the cell

			if state == 0 && liveNeighbours == 3 { //	if the cell is dead and has 3 live neighbours...
				nextGen.matrix[c][r] = 1 //	...revive it
			} else if state == 1 && (liveNeighbours < 2 || liveNeighbours > 3) { //	if cell is alive in hostile conditions...
				nextGen.matrix[c][r] = 0 //	...kill it
			} else { //	otherwise...
				nextGen.matrix[c][r] = g.matrix[c][r] //	...no change in state
			}

		}
	}

	//	Move to the next generation
	g.matrix = nextGen.matrix
}

//	Render the grid onto the screen
func (g *Grid) render(char string) {

	os.Stdout.Write([]byte(ClearScreen))

	str := ""
	for c := 0; c < g.columns; c++ {
		for r := 0; r < g.rows; r++ {
			if g.matrix[c][r] == 1 {
				str += fmt.Sprintf("\u001b[%d;%dH", r, c)
				str += char
			}
		}
	}

	os.Stdout.Write([]byte(str)) //	Print the grid

}
