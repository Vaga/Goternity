package main

import (
	"flag"
	"github.com/vaga/goternity/population"
	"math/rand"
	//"runtime"
	"fmt"
	"time"
)

var input = flag.String("in", "new", "Input file : [name].goternity")
var output = flag.String("out", "result.goternity", "Output file : [name].goternity")
var render = flag.String("render", "render.png", "Render file : [name].png")
var loop = flag.Int("loop", 100, "Loop repetition : 100")
var populationSize = flag.Int("pop", 100, "Population size : 100")

func main() {

	flag.Parse()

	//runtime.GOMAXPROCS(runtime.NumCPU())

	// Init the random seed
	rand.Seed(time.Now().Unix())

	// 1 - Create or load a population (Default: 100)
	var pop *population.Population
	var err error

	if *input == "new" {
		pop = population.New(*populationSize)
	} else {
		pop, err = population.Load(*input)
		if err != nil {
			panic(err)
		}
	}

	// 2 - First evaluation
	pop.Evaluation()

	bestScore := 0.0

	// 3 - Main loop
	for i := 0; i < *loop; i++ {

		// !Reproduction

		// Selection tournament
		pop.Selection()
		//  |->Crossover Swap Region => 0.9
		pop.Crossover()

		// Mutation Rotation/Swap => 0.1
		pop.Mutation()

		// Clean
		pop.Clean()

		// Evaluation
		pop.Evaluation()

		// Information
		pop.Generation++
		pop.Info()

		// Elitism
		pop.Elitism()

		// Draw the best result
		if bestScore < pop.Elite.Score {
			pop.Elite.Render(*render)
			bestScore = pop.Elite.Score
		}

		time.Sleep(20 * time.Millisecond)
	}

	fmt.Println("Final Score : ", bestScore)

	// Save the population
	pop.Save(*output)
}
