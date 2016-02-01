package board

import (
	"bufio"
	"fmt"
	"github.com/vaga/goternity/puzzle"
	"image"
	"image/draw"
	"image/png"
	"io"
	"os"
)

const BOARD_SIZE = 16

type Board struct {
	Pieces [][]*puzzle.Piece
	Score  float64
}

func New() (*Board, error) {

	board := &Board{
		Pieces: make([][]*puzzle.Piece, BOARD_SIZE),
		Score:  0,
	}

	// Generate new puzzle and shuffle it
	p := puzzle.New()
	p.Shuffle()

	for y := 0; y < BOARD_SIZE; y++ {

		board.Pieces[y] = make([]*puzzle.Piece, BOARD_SIZE)
		for x := 0; x < BOARD_SIZE; x++ {
			board.Pieces[y][x] = p[y*BOARD_SIZE+x]
		}
	}

	return board, nil
}

func Load(scanner *bufio.Scanner) (*Board, error) {

	board := &Board{
		Score:  0,
		Pieces: make([][]*puzzle.Piece, BOARD_SIZE),
	}

	for y := 0; y < BOARD_SIZE; y++ {

		board.Pieces[y] = make([]*puzzle.Piece, BOARD_SIZE)
		for x := 0; x < BOARD_SIZE; x++ {

			var id, orientation int
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d %d", &id, &orientation)

			board.Pieces[y][x] = new(puzzle.Piece)
			*board.Pieces[y][x] = *puzzle.DefaultPuzzle[id-1]
			board.Pieces[y][x].Orientation = orientation
		}
	}

	return board, nil
}

func (b *Board) Save(file *os.File) error {

	for y := 0; y < BOARD_SIZE; y++ {
		for x := 0; x < BOARD_SIZE; x++ {
			if _, err := io.WriteString(file, fmt.Sprintf("%d %d\n", b.Pieces[y][x].Id, b.Pieces[y][x].Orientation)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *Board) Render(filename string) error {

	dst := image.NewRGBA(image.Rect(0, 0, 155*16, 155*16))

	for y, col := range b.Pieces {
		for x, piece := range col {
			img := piece.Render()
			draw.Draw(dst, image.Rect(x*155, y*155, 155*16, 155*16), img, image.ZP, draw.Src)
		}
	}

	// TODO: goroutine ?
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	png.Encode(file, dst)

	return nil
}

func (b Board) Debug() {

	for _, col := range b.Pieces {
		for _, piece := range col {

			fmt.Printf("%f ", piece.Score)
		}
		fmt.Println()
	}
	fmt.Printf("Score : %f/%f\n", b.Score, float64(BOARD_SIZE*BOARD_SIZE))
}
