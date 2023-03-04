package singleton

type ISingleton[T any] interface {
	GetInstance() T
}

type Service struct {
	counter int
}

func newService() *Service {
	return &Service{
		counter: 0,
	}
}

func (s *Service) Increment() {
	s.counter++
}

type Singleton struct {
	Service *Service
}

func (s *Singleton) GetInstance() *Service {
	if s.Service == nil {
		s.Service = newService()
	}
	return s.Service
}
