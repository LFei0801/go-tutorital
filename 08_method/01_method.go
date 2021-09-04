package main

type triangle struct {
	size int
}

type colorTriangle struct {
	triangle
	color string
}

type square struct {
	size int
}

// 计算三角形周长
func (t triangle) perimeter() int{
	return t.size * 3
}

// 重载周长方法
func (c colorTriangle) perimeter() int{
	return c.size * 3 * 2
}

func (t *triangle) doubleSize(){
	t.size *= 2
}

//计算正方形周长
func (s square) perimeter() int{
	return s.size * 4
}



