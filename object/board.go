package object

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

type Board [][]*Piece

func NewBoard(size int) Board {

	ShufflePieces()

	board := make(Board, size)

	for y := 0; y < size; y++ {

		board[y] = make([]*Piece, size)
		for x := 0; x < size; x++ {
			board[y][x] = Pieces[y*size+x]
		}
	}

	return board
}

func (b Board) Evaluate() {

	for y, col := range b {
		for x, piece := range col {

			// North
			if (y-1 < 0 && piece.North() == 0) ||
				(y-1 >= 0 && b[y-1][x].South() == piece.North()) {
				piece.Score = piece.Score + 0.25
			}
			//East
			if (x+1 >= 16 && piece.East() == 0) ||
				(x+1 < 16 && b[y][x+1].West() == piece.East()) {
				piece.Score = piece.Score + 0.25
			}
			//South
			if (y+1 >= 16 && piece.South() == 0) ||
				(y+1 < 16 && b[y+1][x].North() == piece.South()) {
				piece.Score = piece.Score + 0.25
			}
			//West
			if (x-1 < 0 && piece.West() == 0) ||
				(x-1 >= 0 && b[y][x-1].East() == piece.West()) {
				piece.Score = piece.Score + 0.25
			}
		}
	}
	b.Debug()
}

func (b *Board) Render() {

	dst := image.NewRGBA(image.Rect(0, 0, 155*16, 155*16))

	for y, col := range *b {
		for x, piece := range col {
			img := piece.Render()
			draw.Draw(dst, image.Rect(x*155, y*155, 155*16, 155*16), img, image.ZP, draw.Src)
		}
	}

	file, err := os.Create("./result.png")
	if err != nil {
		fmt.Println("Ooppps... I can't create result.png")
	}
	defer file.Close()

	png.Encode(file, dst)
}
func (b Board) Debug() {

	for _, col := range b {
		for _, piece := range col {

			fmt.Printf("%f ", piece.Score)
		}
		fmt.Println()
	}
}
