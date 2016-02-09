package main

import (
	"flag"
	//"fmt"
	"github.com/vaga/goternity/population"
	"math/rand"
	//"runtime"
	"time"
)

var input = flag.String("in", "new", "Input file : [name].goternity")
var output = flag.String("out", "result.goternity", "Output file : [name].goternity")
var render = flag.String("render", "render.png", "Render file : [name].png")

func main() {

	flag.Parse()

	//runtime.GOMAXPROCS(runtime.NumCPU())

	// Init the random seed
	rand.Seed(time.Now().Unix())

	// 1 - Create or load a population (Default: 100)
	var pop *population.Population
	var err error

	if *input == "new" {
		pop = population.New(100)
	} else {
		pop, err = population.Load(*input)
		if err != nil {
			panic(err)
		}
	}

	// 2 - First evaluation
	pop.Evaluation()

	// 3 - Main loop
	for i := 0; i < 1000; i++ {

		// !Reproduction

		// Selection tournament
		pop.Selection()
		//  |->Crossover Echange de region(1st and 2nd) 0.9
		pop.Crossover()

		// Mutation Rotation/Swap() 0.1
		pop.Mutation()

		// Clean
		pop.Clean()

		// Evaluation
		pop.Evaluation()
		pop.Generation++
		pop.Info()

		// Elitism
		pop.Elitism()

		time.Sleep(20 * time.Millisecond)
	}

	// Save the population
	pop.Save(*output)
}
