package object

type Square struct {
	Sisi int
}

func (s Square) GetResultA() int {
	return s.Sisi * 4
}

func (s Square) GetResultB() int {
	return s.Sisi / 4
}
