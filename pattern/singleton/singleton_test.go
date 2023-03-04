package singleton

import "testing"

func TestSingleton(t *testing.T) {
	singleton := Singleton{}

	if singleton.Service != nil {
		t.Errorf("Service %v; want nil have %v", singleton.Service, singleton.Service)
	}

	service := singleton.GetInstance()

	if service == nil {
		t.Errorf("Service %v; want not nil have %v", service, service)
	}
}
func TestIncrement(t *testing.T) {

	var singleton Singleton

	service := singleton.GetInstance()

	service.Increment()

	service2 := singleton.GetInstance()

	service2.Increment()

	if service.counter != 2 {
		t.Errorf("Counter %d; want 2 have %d", service.counter, service.counter)
	}

	service.Increment()
	service.Increment()

	if service2.counter != 4 {
		t.Errorf("Counter %d; want 4 have %d", service2.counter, service2.counter)
	}
}
