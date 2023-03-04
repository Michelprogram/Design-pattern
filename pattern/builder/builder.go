package builder

type Builder interface {
	Seats(int) Builder
	Wheels(int) Builder
	Mark(string) Builder
	GetResult() Builder
	Reset() Builder
}

type CarBuilder struct {
	Seats_nb  int
	Wheels_nb int
	Mark_nb   string
}

var _ Builder = (*CarBuilder)(nil)

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (cb *CarBuilder) Seats(nb int) Builder {
	cb.Seats_nb = nb
	return cb
}

func (cb *CarBuilder) Wheels(nb int) Builder {
	cb.Wheels_nb = nb
	return cb
}

func (cb *CarBuilder) Mark(mark string) Builder {
	cb.Mark_nb = mark
	return cb
}

func (cb *CarBuilder) GetResult() Builder {
	return cb
}

func (cb *CarBuilder) Reset() Builder {
	return NewCarBuilder()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) ConstructBerlin() Builder {
	return d.builder.Seats(4).Wheels(4).Mark("BMW")
}

func (d *Director) ConstructParis() Builder {
	return d.builder.Seats(2).Wheels(2).Mark("Peugeot")
}
