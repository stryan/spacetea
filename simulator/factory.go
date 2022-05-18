package simulator

//Factory is a game object that can hold resources and produce something
type Factory struct {
	Resources   map[int]int
	Description string
	Formula     string
}

func (f *Factory) Tick() {
	panic("not implemented") // TODO: Implement
}

func (f *Factory) Get() Produce {
	panic("not implemented") // TODO: Implement
}

func (f *Factory) String() string {
	panic("not implemented") // TODO: Implement
}
