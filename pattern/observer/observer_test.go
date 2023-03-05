package observer

import "testing"

func TestSubscriber(t *testing.T) {
	j := NewJournal()

	mobile := &Mobile{}

	j.Subscribe(&Mobile{})
	j.Subscribe(mobile)
	j.Subscribe(&Mail{})

	if len(j.GetSubscribers()) != 3 {
		t.Error("Expected 3 subscribers")
	}

	j.Unsubscribe(mobile)

	if len(j.GetSubscribers()) != 2 {
		t.Error("Expected 2 subscribers")
	}
}

func TestNotifyAll(t *testing.T) {
	j := NewJournal()

	mobile := &Mobile{}

	j.Subscribe(&Mobile{})
	j.Subscribe(mobile)
	j.Subscribe(&Mail{})

	j.NotifyAll("Hello World")

	if j.GetNotified() != 3 {
		t.Error("Expected 3 subscribers")
	}
}
