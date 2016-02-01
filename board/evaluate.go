package board

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

	return b.Score
}
