package board

type Board struct {
	Size    int
	Snakes  map[int]int
	Ladders map[int]int
}

func NewBoard(size int, snakes []*Snake, ladders []*Ladder) *Board {
	board := &Board{
		Size:    size,
		Snakes:  make(map[int]int),
		Ladders: make(map[int]int),
	}
	for _, snake := range snakes {
		board.Snakes[snake.Head] = snake.Tail
	}
	for _, ladder := range ladders {
		board.Ladders[ladder.Bottom] = ladder.Top
	}
	return board
}

func (b *Board) GetNewPosition(position int) int {
	if tail, ok := b.Snakes[position]; ok {
		return tail
	}
	if top, ok := b.Ladders[position]; ok {
		return top
	}
	return position
}
