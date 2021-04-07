package model

type Segment [][]int

func NewSegment(x, y int) Segment {
	s := Segment(make([][]int, x))

	for col := range s {
		s[col] = make([]int, y)
	}

	return s
}
