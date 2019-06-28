package object

type Circle struct {
	Diameter int
}

func (c Circle) GetJariJari() int {
	return c.Diameter / 2
}

func (c Circle) GetResultA() int {
	return c.Diameter * 2
}

func (c Circle) GetResultB() int {
	return (c.Diameter + 2) * 2
}
