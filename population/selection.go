package population

import (
	"github.com/vaga/goternity/board"
	"math/rand"
)

const (
	SELECTION_NB_INDIVIDUAL = 4
	SELECTION_POOL_SIZE     = 10
)

// Tournament Selection
func (p *Population) Selection() {

	p.SelectedBoards = make([]*board.Board, SELECTION_NB_INDIVIDUAL)

	for i := 0; i < SELECTION_NB_INDIVIDUAL; i++ {

		var best *board.Board = nil

		for k := 0; k < SELECTION_POOL_SIZE; k++ {

			current := p.Boards[rand.Intn(len(p.Boards))]

			if best == nil || current.Score > best.Score {
				best = current
			}
		}
		p.SelectedBoards[i] = best
	}
}
