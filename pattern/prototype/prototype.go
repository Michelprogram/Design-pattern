package prototype

type Prototype interface {
	Clone() Prototype
}

type Shapes struct {
	x     int
	y     int
	color string
}

var _ Prototype = (*Shapes)(nil)

func NewShapes(x, y int, color string) *Shapes {
	return &Shapes{x: x, y: y, color: color}
}

func (s *Shapes) Clone() Prototype {
	return &Shapes{x: s.x, y: s.y, color: s.color}
}

type Rectangle struct {
	width int
	height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func (r *Rectangle) Clone() Prototype {
	return &Rectangle{width: r.width, height: r.height}
}

type Circle struct {
	radius int
}

func NewCircle(radius int) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) Clone() Prototype {
	return &Circle{radius: c.radius}
}