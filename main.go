package main

import (
	//"fmt"
	"github.com/vaga/goternity/object"
	//"time"
	"flag"
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

	// 2 - First evaluation
	board.Evaluate()

	// TODO:
	// 3 - Main loop
	// for !board.Done() {

	//	fmt.Println("TODO:\n - Selection;\n - Reproduction;\n - Crossover;\n - Mutation;\n - Evaluation;\n")
	//	time.Sleep(100 * time.Millisecond)
	//}

	if err := board.Render(*render); err != nil {
		panic(err)
	}
}
