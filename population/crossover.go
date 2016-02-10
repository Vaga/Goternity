package population

import (
	"fmt"
	"github.com/vaga/goternity/board"
	"math/rand"
)

/**
 * Crossover : Region Exchange
 * Rate : 0.9
 * Range : [2; 5]
 */
func (p *Population) Crossover() {

	for i := 0; i < len(p.SelectedBoards); i += 2 {

		if rand.Intn(100) < 90 {
			fmt.Println("Crossover : Region Exchange")
			doCrossover(p.SelectedBoards[i], p.SelectedBoards[i+1])
			p.NbCrossover++
		}
	}
}

func doCrossover(board1, board2 *board.Board) {

	width := 2 + rand.Intn(4)
	height := 2 + rand.Intn(4)
	startX := rand.Intn(board.BOARD_SIZE - width)
	startY := rand.Intn(board.BOARD_SIZE - height)

	for x := startX; x <= width; x++ {
		for y := startY; y <= height; y++ {
			board1.Pieces[y][x], board2.Pieces[y][x] = board2.Pieces[y][x], board1.Pieces[y][x]
		}
	}
}
