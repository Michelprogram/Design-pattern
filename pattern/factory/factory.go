package factory

type Button interface {
	Click() string
}

type IOSButton struct {
}

func (b *IOSButton) Click() string {
	return "IOS button clicked"
}

type AndroidButton struct {
}

func (b *AndroidButton) Click() string {
	return "Android button clicked"
}

type Factory struct{
}

func (f *Factory) CreateButton(os string) Button {
	if os == "ios" {
		return &IOSButton{}
	}
	return &AndroidButton{}
}