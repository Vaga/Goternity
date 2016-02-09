package population

import (
	"fmt"
	"github.com/vaga/goternity/board"
	"math/rand"
)

func (p *Population) Mutation() {

	for i := 0; i < len.(p.SelectedBoards); i++ {
		
		if rand.Intn(100) < 10 {

			fmt.Println("Mutation: Region Rotation and Swap")

			doMutation(p.SelectedBoards[i])
		}
	}
}

func doMutation(board) {
	
	x := rand.Intn(board.BOARD_SIZE - 2)
	y := rand.Intn(board.BOARD_SIZE - 2)
	x2 := rand.Intn(board.BOARD_SIZE - 2)
	y2 := rand.Intn(board.BOARD_SIZE - 2)

	un, deux, trois, quatre = board.Piece[x][y], board.Piece[x + 1][y], board.Piece[x + 1][y + 1], board.Piece[x][y + 1]

	one, two, three, four = board.Piece[x2][y2], board.Piece[x2 + 1][y2], board.Piece[x2 + 1][y2 + 1], board.Piece[x2][y2 + 1]

	board.Piece[x][y], board.Piece[x + 1][y], board.Piece[x + 1][y + 1], board.Piece[x][y + 1] = four, one, two, three

	board.Piece[x2][y2], board.Piece[x2 + 1][y2], board.Piece[x2 + 1][y2 + 1], board.Piece[x2][y2 + 1] = quatre, un, deux, trois

}