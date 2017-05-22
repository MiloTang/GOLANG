package test11

type Rect struct {
	x, y int
}

func (r *Rect) Area() int {
	return r.x * r.y
}
func NewRect(x, y int) *Rect {
	return &Rect{x, y}
}

func Show() {
	r := &Rect{2, 3}
	r.Area()
}
