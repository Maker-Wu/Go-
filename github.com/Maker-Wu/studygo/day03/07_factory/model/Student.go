package model

type student struct {
	Name string
	score float64
}

func (s *student) GetScore() float64 {
	return s.score
}

// NewStudent 因为student结构体首字母是小写，因此只能在model使用
func NewStudent(n string, s float64) *student {
	return &student{
		Name: n,
		score: s,
	}
}
