package simulator

//Producer is a game object that producers a material
type Producer interface {
	Tick()
	Get() Produce
	String() string
	Describe() string
}

//Produce is the result of a producer
type Produce struct {
	Kind  int
	Value int
}
