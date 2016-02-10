package population

import (
	"fmt"
	"github.com/vaga/goternity/board"
	"github.com/vaga/goternity/puzzle"
	"math/rand"
)

func (p *Population) Mutation() {

	for i := 0; i < len(p.Boards); i++ {

		if rand.Intn(100) < 10 {
			fmt.Println("Mutation: Region Rotation and Swap")
			doMutation(p.Boards[i])
		}
	}
}

func doMutation(b *board.Board) {

	size := 2 + rand.Intn(4)

	x1 := rand.Intn(board.BOARD_SIZE - size)
	y1 := rand.Intn(board.BOARD_SIZE - size)
	x2 := rand.Intn(board.BOARD_SIZE - size)
	y2 := rand.Intn(board.BOARD_SIZE - size)

	region1 := make([]*puzzle.Piece, 0)
	region2 := make([]*puzzle.Piece, 0)

	for x := 0; x < size; x++ {
		for y := size - 1; y >= 0; y-- {

			// Rotate all pieces of 90 deg
			b.Pieces[y1+y][x1+x].Rotate(1)
			b.Pieces[y2+y][x2+x].Rotate(1)

			// Save pieces in a temp list
			region1 = append(region1, b.Pieces[y1+y][x1+x])
			region2 = append(region2, b.Pieces[y2+y][x2+x])
		}
	}

	// Replace the pieces
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			b.Pieces[y1+y][x1+x] = region2[y*size+x]
			b.Pieces[y2+y][x2+x] = region1[y*size+x]
		}
	}
}
