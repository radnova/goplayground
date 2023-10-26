package main

import "fmt"

const numRows = 11
const numCols = 11

var board [numRows][numCols]int

func main() {
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			board[i][j] = 0
		}
	}
	colHeader := "A"
	fmt.Print(" ")
	for i := 1; i < numCols; i++ {
		fmt.Printf("%s ", colHeader)
		nextRune := []rune(colHeader)[0] + 1
		colHeader = string(nextRune)
	}
	fmt.Println("")

	for i:= 0; i < numCols; i++ {
		for j := 0
	}
}
