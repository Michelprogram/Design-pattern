package factory

import "testing"

func TestFactory(t *testing.T) {
	f := &Factory{}
	iosButton := f.CreateButton("ios")
	androidButton := f.CreateButton("android")
	if iosButton.Click() != "IOS button clicked" {
		t.Errorf("Expected IOS button clicked, got %s", iosButton.Click())
	}
	if androidButton.Click() != "Android button clicked" {
		t.Errorf("Expected Android button clicked, got %s", androidButton.Click())
	}
}
