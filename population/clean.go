package population

import (
	"fmt"
	"github.com/vaga/goternity/board"
	"github.com/vaga/goternity/puzzle"
	"math/rand"
)

func (p *Population) Clean() {


	for a := 0; a < len.(p.SelectedBoard); a++ {
		
		for i := 0; i < BOARD_SIZE * BOARD_SIZE, i++ {

			for j := 0; j < BOARD_SIZE * BOARD_SIZE, j++ {
	
			//  IF   p.SelectedBoard[a].Piece[i % BOARD_SIZE][i / BOARD_SIZE] == p.SelectedBoard[a].Piece[j % BOARD_SIZE][j / BOARD_SIZE]
			// AND I != J

			}
			j = 0;

		}
		i = 0;
	}


	for i := 0; i < len(p.Boards); i++ {

		if rand.Intn(100) < 10 {
			fmt.Println("Clean Population")
			doClean(p.Boards[i])
			p.NbMutation++
		}
	}


}

func doClean (b *board.Board) {

	in := list.New()
	out := list.New()
	missing = 1;

	for i := 0; i < 256; i++ {
		
		for x := 0; x < BOARD_SIZE; x++ {
		
			for y = 0; y < BOARD_SIZE; y++ {
				if (b.Pieces[x][y].Id == i)
					missing = 0;
			}
			y = 0;
		}
		x = 0; 
		if (missing == 0) {
			missing = 1
			e := in.PushBack(i);
		}
	}

	for i := 0; i < BOARD_SIZE * BOARD_SIZE, i++ {

		for j := 0; j < BOARD_SIZE * BOARD_SIZE, j++ {
	
			if (i != j) {

				if (b.Pieces[i % BOARD_SIZE][i / BOARD_SIZE].Id == b.Pieces[j % BOARD_SIZE][j / BOARD_SIZE].Id) {
					
					e = in.Back()
					newPiece := new(Piece)
					*newPiece = *DefaultPuzzle[in.Back()]

					b.Pieces[i % BOARD_SIZE][i / BOARD_SIZE] = newPiece

					in.Remove(e)
				}
			}
		}
		j = 0;

	}

	 
}
