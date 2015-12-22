package object

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

const BOARD_SIZE = 16

type Board struct {
	Pieces [][]*Piece
	Score  float64
}

func NewBoard() (*Board, error) {

	board := &Board{
		Pieces: make([][]*Piece, BOARD_SIZE),
		Score:  0,
	}

	puzzle := NewPuzzle()
	puzzle.Shuffle()

	for y := 0; y < BOARD_SIZE; y++ {

		board.Pieces[y] = make([]*Piece, BOARD_SIZE)
		for x := 0; x < BOARD_SIZE; x++ {
			board.Pieces[y][x] = puzzle[y*BOARD_SIZE+x]
		}
	}

	return board, nil
}

func LoadBoard(filename string) (*Board, error) {

	board := &Board{
		Score:  0,
		Pieces: make([][]*Piece, BOARD_SIZE),
	}

	puzzle, err := LoadPuzzle(filename)
	if err != nil {
		return nil, err
	}

	for y := 0; y < BOARD_SIZE; y++ {

		board.Pieces[y] = make([]*Piece, BOARD_SIZE)
		for x := 0; x < BOARD_SIZE; x++ {
			board.Pieces[y][x] = puzzle[y*BOARD_SIZE+x]
		}
	}

	return board, nil
}

func (b *Board) Evaluate() {

	b.Score = 0

	for y, col := range b.Pieces {
		for x, piece := range col {

			// TODO: If 2 edges == 0

			// North
			if (y-1 < 0 && piece.North() == 0) ||
				(y-1 >= 0 && b.Pieces[y-1][x].South() == piece.North()) {
				piece.Score = piece.Score + 0.25
			}
			//East
			if (x+1 >= BOARD_SIZE && piece.East() == 0) ||
				(x+1 < BOARD_SIZE && b.Pieces[y][x+1].West() == piece.East()) {
				piece.Score = piece.Score + 0.25
			}
			//South
			if (y+1 >= BOARD_SIZE && piece.South() == 0) ||
				(y+1 < BOARD_SIZE && b.Pieces[y+1][x].North() == piece.South()) {
				piece.Score = piece.Score + 0.25
			}
			//West
			if (x-1 < 0 && piece.West() == 0) ||
				(x-1 >= 0 && b.Pieces[y][x-1].East() == piece.West()) {
				piece.Score = piece.Score + 0.25
			}
			b.Score = b.Score + piece.Score
		}
	}
	b.Debug()
}

func (b *Board) Done() bool {
	return (int(b.Score) == BOARD_SIZE*BOARD_SIZE)
}

func (b *Board) Save(filename string) error {

	puzzle := new(Puzzle)

	for y := 0; y < BOARD_SIZE; y++ {
		for x := 0; x < BOARD_SIZE; x++ {
			puzzle[y*BOARD_SIZE+x] = b.Pieces[y][x]
		}
	}

	if err := puzzle.Save(filename); err != nil {
		return err
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
	fmt.Printf("Score : %f\n", b.Score)
}
