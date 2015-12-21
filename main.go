package main

import (
	"github.com/vaga/goternity/object"
)

func main() {

	board := object.NewBoard(16)
	board.Evaluate()

	board.Render()
}
