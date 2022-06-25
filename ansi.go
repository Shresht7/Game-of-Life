package main

import "fmt"

//	Escape Code
const ESC = "\u001b"

//	Control Sequence Introducer
const CSI = ESC + "["

//	Clear out the entire screen
const ClearScreen = CSI + "2J"

//	Move cursor to the given row and column
func CursorMoveTo(row, column int) string {
	return fmt.Sprintf("%s%d;%dH", CSI, row, column)
}
