package object

import (
	"bufio"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/png"
	"math/rand"
	"os"
)

type Piece struct {
	Id          int
	south       int
	east        int
	north       int
	west        int
	Orientation int
	image       image.Image
	Score       float64
}

var Pieces [256]*Piece

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
		Pieces[i] = piece
	}
}

func ShufflePieces() {

	rand.Seed(42)

	for i := range Pieces {
		j := rand.Intn(i + 1)
		Pieces[i].Orientation = rand.Intn(4)
		Pieces[i], Pieces[j] = Pieces[j], Pieces[i]
	}
}

func (p *Piece) North() int {
	if p.Orientation == 1 {
		return p.west
	} else if p.Orientation == 2 {
		return p.south
	} else if p.Orientation == 3 {
		return p.east
	}
	return p.north
}

func (p *Piece) East() int {
	if p.Orientation == 1 {
		return p.north
	} else if p.Orientation == 2 {
		return p.west
	} else if p.Orientation == 3 {
		return p.south
	}
	return p.east
}

func (p *Piece) South() int {
	if p.Orientation == 1 {
		return p.east
	} else if p.Orientation == 2 {
		return p.north
	} else if p.Orientation == 3 {
		return p.west
	}
	return p.south
}

func (p *Piece) West() int {
	if p.Orientation == 1 {
		return p.south
	} else if p.Orientation == 2 {
		return p.east
	} else if p.Orientation == 3 {
		return p.north
	}
	return p.west
}

func (p *Piece) Render() image.Image {

	if p.Orientation == 1 {
		return imaging.Rotate270(p.image)
	} else if p.Orientation == 2 {
		return imaging.Rotate180(p.image)
	} else if p.Orientation == 3 {
		return imaging.Rotate90(p.image)
	}
	return p.image
}
