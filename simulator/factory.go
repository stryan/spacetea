package simulator

//Factory is a game object that can hold resources and produce something
type Factory struct {
	Resources   map[int]int
	value       int
	outputType  int
	Description string
	formula     string
}

//Tick a beat
func (f *Factory) Tick() {
	panic("not implemented") // TODO: Implement
}

//Get produce
func (f *Factory) Get() Produce {
	panic("not implemented") // TODO: Implement
}

//String
func (f *Factory) String() string {
	panic("not implemented") // TODO: Implement
}
