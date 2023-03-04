package builder

import "testing"

func TestBuilder(t *testing.T) {
	
	builder := NewCarBuilder()

	car := builder.Seats(4).Wheels(4).Mark("Tesla").GetResult()

	if car == nil {
		t.Error("Car is nil")
	}

	if builder.Seats_nb != 4 {
		t.Errorf("Expected 4 seats, got %d", builder.Seats_nb)
	}
}

func TestDirector(t *testing.T){
	builder := NewCarBuilder()
	director := NewDirector(builder)

	car := director.ConstructBerlin().GetResult()

	if car == nil {
		t.Error("Car is nil")
	}

	if builder.Seats_nb != 4 {
		t.Errorf("Expected 4 seats, got %d", builder.Seats_nb)
	}

	if builder.Wheels_nb != 4 {
		t.Errorf("Expected 4 wheels, got %d", builder.Wheels_nb)
	}

	if builder.Mark_nb != "BMW" {
		t.Errorf("Expected BMW mark, got %s", builder.Mark_nb)
	}
}