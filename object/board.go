package object

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math/rand"
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

func (b *Board) Evaluate() float64 {

	b.Score = 0

	for y, col := range b.Pieces {
		for x, piece := range col {

			piece.Score = 0

			// North
			if (y-1 < 0 && piece.North() == 0) ||
				(y-1 >= 0 && b.Pieces[y-1][x].South() == piece.North() && piece.North() != 0) {
				piece.Score = piece.Score + 0.25
			}
			//East
			if (x+1 >= BOARD_SIZE && piece.East() == 0) ||
				(x+1 < BOARD_SIZE && b.Pieces[y][x+1].West() == piece.East() && piece.East() != 0) {
				piece.Score = piece.Score + 0.25
			}
			//South
			if (y+1 >= BOARD_SIZE && piece.South() == 0) ||
				(y+1 < BOARD_SIZE && b.Pieces[y+1][x].North() == piece.South() && piece.South() != 0) {
				piece.Score = piece.Score + 0.25
			}
			//West
			if (x-1 < 0 && piece.West() == 0) ||
				(x-1 >= 0 && b.Pieces[y][x-1].East() == piece.West() && piece.West() != 0) {
				piece.Score = piece.Score + 0.25
			}
			b.Score = b.Score + piece.Score
		}
	}
	b.Debug()

	return b.Score
}

func (b *Board) Random() {

	for i := 0; i < 50; i++ {

		y1 := rand.Intn(BOARD_SIZE)
		x1 := rand.Intn(BOARD_SIZE)
		y2 := rand.Intn(BOARD_SIZE)
		x2 := rand.Intn(BOARD_SIZE)
		piece1 := b.Pieces[y1][x1]
		piece2 := b.Pieces[y2][x2]

		maxPiece := 0.0
		if piece1.Score > piece2.Score {
			maxPiece = piece1.Score
		} else {
			maxPiece = piece2.Score
		}

		prob := 1 - maxPiece
		de := float64(rand.Intn(100)) / 100

		if de < prob {

			if de < 0.50 {
				b.Pieces[y1][x1].Rotate(rand.Intn(3) + 1)
			}

			b.Pieces[y1][x1], b.Pieces[y2][x2] = b.Pieces[y2][x2], b.Pieces[y1][x1]
		}
	}
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
	fmt.Printf("Score : %f/%f\n", b.Score, float64(BOARD_SIZE*BOARD_SIZE))
}
