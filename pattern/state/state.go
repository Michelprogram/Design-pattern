package state

type State interface {
	Render() string
	Publish() string
	ChangeDocument(d *Document)
}

type Document struct {
	state State
	Name  string
}

func NewDocument(name string) *Document {
	return &Document{Name: name}
}

func (d *Document) Render() {
	d.state.Render()
}

func (d *Document) Publish() {
	d.state.Publish()
}

func (d *Document) ChangeState(state State) {
	d.state = state
}

type Brouillon struct {
	document *Document
}

func (b *Brouillon) Render() string {
	return "Render in brouillon"
}

func (b *Brouillon) Publish() string {
	return "Publish in brouillon"
}

func (b *Brouillon) ChangeDocument(document *Document) {
	b.document = document
}

type EnCours struct {
	document *Document
}

func (e *EnCours) Render() string {
	return "Render in en cours"
}

func (e *EnCours) Publish() string {
	return "Publish in en cours"
}

func (e *EnCours) ChangeDocument(document *Document) {
	e.document = document
}

type Publie struct {
	document *Document
}

func (p *Publie) Render() string {
	return "Render in publie"
}

func (p *Publie) Publish() string {
	return "Publish in publie"
}

func (p *Publie) ChangeDocument(document *Document) {
	p.document = document
}
