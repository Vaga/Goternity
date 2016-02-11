package population

import (
	"fmt"
	"github.com/vaga/goternity/board"
	"github.com/vaga/goternity/puzzle"
	//"math/rand"
)

func (p *Population) Clean() {

	for i := 0; i < len(p.Boards); i++ {

		fmt.Println("Clean Population")
		doClean(p.Boards[i])
		p.NbClean++
	}

}

func doClean(b *board.Board) {

	missing := make(map[int]int)
	for i := 1; i <= board.BOARD_SIZE*board.BOARD_SIZE; i++ {
		missing[i] = i
	}
	unique := make(map[int]int)
	double := make(map[int]int)

	for x := 0; x < board.BOARD_SIZE; x++ {

		for y := 0; y < board.BOARD_SIZE; y++ {

			p := b.Pieces[y][x]

			// Double ?
			if _, ok := unique[p.Id]; ok {
				delete(unique, p.Id)
				double[p.Id] = p.Id
				continue
			}

			// Unique ?
			if _, ok := missing[p.Id]; ok {
				delete(missing, p.Id)
				unique[p.Id] = p.Id
				continue
			}

		}
	}

	for x := 0; x < board.BOARD_SIZE; x++ {

		for y := 0; y < board.BOARD_SIZE; y++ {
			p := b.Pieces[y][x]

			if _, ok := double[p.Id]; ok {
				for id := range missing {
					*b.Pieces[y][x] = *puzzle.DefaultPuzzle[id-1]
					delete(missing, id)
					break
				}
			}
		}
	}
}
