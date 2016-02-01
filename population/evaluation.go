package population

func (p *Population) Evaluation() {

	p.Score = 0.0
	for _, board := range p.Boards {
		board.Evaluate()
		if board.Score > p.Score {
			p.Score = board.Score
		}
	}
}
