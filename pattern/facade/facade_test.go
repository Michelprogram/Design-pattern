package facade

import "testing"

func TestFacade(t *testing.T) {
	facade := NewOrderFacade()

	if facade.Order() != 1 {
		t.Error("Order() should return 1")
	}

}
