package main

import (
	"flag"
	"fmt"
	"github.com/vaga/goternity/object"
	"math/rand"
	"time"
)

var input = flag.String("in", "new", "Input file : [name].goternity")
var output = flag.String("out", "result.goternity", "Output file : [name].goternity")
var render = flag.String("render", "render.png", "Render file : [name].png")

func main() {

	flag.Parse()

	var board *object.Board
	var err error

	// 1 - Create a population (Board)
	if *input == "new" {
		board, _ = object.NewBoard()
		if err = board.Save(*output); err != nil {
			panic(err)
		}
	} else {
		board, err = object.LoadBoard(*input)
		if err != nil {
			panic(err)
		}
	}

	rand.Seed(time.Now().Unix())

	// 2 - First evaluation
	bestScore := board.Evaluate()
	try := 0

	// 3 - Main loop
	for !board.Done() {

		fmt.Printf("#%d ----------\n", try)

		board.Random()
		currentScore := board.Evaluate()

		if currentScore > bestScore {

			board.Save(*output)
			if err := board.Render(*render); err != nil {
				panic(err)
			}
			bestScore = currentScore
		}
		fmt.Println("-------------")
		time.Sleep(20 * time.Millisecond)
	}
}
