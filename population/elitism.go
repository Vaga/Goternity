package population

import (
	"github.com/vaga/goternity/board"
)

func (p *Population) Elitism() {

	var worst *board.Board = nil
	var worstId int
	var best *board.Board = nil

	for k := 0; k < len(p.Boards); k++ {

		current := p.Boards[k]

		if best == nil || current.Score > best.Score {
			best = current
		}
		if worst == nil || current.Score < worst.Score {
			worst = current
			worstId = k
		}
	}

	if p.Elite != nil {
		p.Boards[worstId] = p.Elite
	}
	p.Elite = best
}
