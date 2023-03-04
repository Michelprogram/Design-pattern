package prototype

import "testing"

func TestPrototype(t *testing.T) {
	rectangle := NewRectangle(10, 20)
	circle := NewCircle(5)
	rectangle2 := rectangle.Clone().(*Rectangle)
	circle2 := circle.Clone().(*Circle)

	if rectangle2.width != rectangle.width {
		t.Errorf("Expected %d, got %d", rectangle.width, rectangle2.width)
	}
	if rectangle2.height != rectangle.height {
		t.Errorf("Expected %d, got %d", rectangle.height, rectangle2.height)
	}
	if circle2.radius != circle.radius {
		t.Errorf("Expected %d, got %d", circle.radius, circle2.radius)
	}
}
