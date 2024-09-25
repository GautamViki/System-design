package board

type Ladder struct {
	Top    int
	Bottom int
}

func NewLadder(top, bottom int) *Ladder {
	return &Ladder{
		Top:    top,
		Bottom: bottom,
	}
}
