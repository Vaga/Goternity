package object

import (
	"bufio"
	"fmt"
	"image/png"
	"io"
	"math/rand"
	"os"
)

const NB_PIECES = 256

type Puzzle [256]*Piece

var defaultPuzzle Puzzle

func init() {

	file, err := os.Open("./assets/pieces.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {

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
		defaultPuzzle[i] = piece
	}
}

func NewPuzzle() *Puzzle {

	newPuzzle := new(Puzzle)

	for i, piece := range defaultPuzzle {
		newPiece := new(Piece)
		*newPiece = *piece
		newPuzzle[i] = newPiece
	}

	return newPuzzle
}

func LoadPuzzle(filename string) (*Puzzle, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	newPuzzle := new(Puzzle)

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {

		var id, orientation int

		fmt.Sscanf(scanner.Text(), "%d %d", &id, &orientation)
		newPuzzle[i] = new(Piece)
		*newPuzzle[i] = *defaultPuzzle[id-1]
		newPuzzle[i].Orientation = orientation
	}

	return newPuzzle, nil
}

func (p *Puzzle) Shuffle() {

	for i := range *p {
		j := rand.Intn(i + 1)
		p[i].Orientation = rand.Intn(4)
		p[i], p[j] = p[j], p[i]
	}
}

func (p Puzzle) Save(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, piece := range p {
		if _, err := io.WriteString(file, fmt.Sprintf("%d %d\n", piece.Id, piece.Orientation)); err != nil {
			return err
		}
	}
	return nil
}
