package simulator

//Object is anything that can be placed in a pod
type Object interface {
	Type() ObjectType
	Tick()
	String() string
	Describe() string
}

//ObjectType is what kind of game object it is
type ObjectType int

const (
	producerObject ObjectType = iota
	consumerObject
)
