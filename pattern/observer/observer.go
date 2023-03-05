package observer

type Journal struct {
	subscribers []Subscriber
	notified    int
	//Variable state pattern
}

func NewJournal() *Journal {
	return &Journal{}
}

func (j *Journal) GetNotified() int {
	return j.notified
}

func (j *Journal) GetSubscribers() []Subscriber {
	return j.subscribers
}

func (j *Journal) Subscribe(s Subscriber) {
	j.subscribers = append(j.subscribers, s)
}

func (j *Journal) Unsubscribe(s Subscriber) {
	for i, v := range j.subscribers {
		if v == s {
			j.subscribers = append(j.subscribers[:i], j.subscribers[i+1:]...)
		}
	}
}

func (j *Journal) NotifyAll(msg string) {
	j.notified = 0
	for _, v := range j.subscribers {
		v.Notify(msg)
		j.notified++
	}
}

type Subscriber interface {
	Notify(string)
}

type Mobile struct {
}

func (m *Mobile) Notify(msg string) {
	println("Mobile: " + msg)
}

type Mail struct {
}

func (m *Mail) Notify(msg string) {
	println("Mail: " + msg)
}
