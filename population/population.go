package population

import (
	"bufio"
	"fmt"
	"github.com/vaga/goternity/board"
	"io"
	"os"
)

type Population struct {
	Generation int
	Boards     []*board.Board
	Score      float64
}

func New(nb int) *Population {

	population := &Population{
		Generation: 0,
		Boards:     make([]*board.Board, nb),
	}

	for i := 0; i < nb; i++ {
		population.Boards[i], _ = board.New()
	}

	return population
}

func Load(filename string) (*Population, error) {

	// Load the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Retrieve the generation and size
	var size, generation int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &generation)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &size)

	// Create the population
	population := &Population{
		Generation: generation,
		Boards:     make([]*board.Board, size),
	}

	// Retrieve the pieces foreach board
	for s := 0; s < size; s++ {
		population.Boards[s], err = board.Load(scanner)
		if err != nil {
			return nil, err
		}
	}

	return population, nil
}

func (p *Population) Info() {

	fmt.Println("--- Goternity ---------")
	fmt.Printf("- Generation  : %d\n", p.Generation)
	fmt.Printf("- len(Boards) : %d\n", len(p.Boards))
	fmt.Printf("- Best Score  : %f\n", p.Score)
	fmt.Println("-----------------------")
}

func (p *Population) Save(filename string) error {

	// Create a file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Add the generation
	if _, err := io.WriteString(file, fmt.Sprintf("%d\n", p.Generation)); err != nil {
		return err
	}

	// Add the size of population
	if _, err := io.WriteString(file, fmt.Sprintf("%d\n", len(p.Boards))); err != nil {
		return err
	}

	// Add the population
	for _, board := range p.Boards {
		if err := board.Save(file); err != nil {
			return err
		}
	}

	return nil
}
