package simulator

//Consumer is a game object that consumes resources
type Consumer interface {
	Object
	Tick()
}
