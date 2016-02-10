package puzzle

import (
	"bufio"
	"fmt"
	"image/png"
	"math/rand"
	"os"
)

const NB_PIECES = 256

type Puzzle [256]*Piece

var DefaultPuzzle Puzzle

func init() {

	file, err := os.Open("./assets/pieces.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {

		func() {
			file, err := os.Open(fmt.Sprintf("./assets/%d.png", i+1))
			if err != nil {
				panic(err)
			}
			defer file.Close()

			img, err := png.Decode(file)
			if err != nil {
				panic(err)
			}

			piece := &Piece{}
			piece.Id = i + 1
			piece.image = img
			fmt.Sscanf(scanner.Text(), "%d %d %d %d", &piece.north, &piece.south, &piece.west, &piece.east)
			DefaultPuzzle[i] = piece
		}()
	}
}

func New() *Puzzle {

	newPuzzle := new(Puzzle)

	for i, piece := range DefaultPuzzle {
		newPiece := new(Piece)
		*newPiece = *piece
		newPuzzle[i] = newPiece
	}

	return newPuzzle
}

func (p *Puzzle) Shuffle() {

	for i := range *p {
		j := rand.Intn(i + 1)
		p[i].Orientation = rand.Intn(4)
		p[i], p[j] = p[j], p[i]
	}
}
