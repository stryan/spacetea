package simulator

import (
	"strconv"
)

//Plant is a plant that grows per tick
type Plant struct {
	kind  int
	value int
}

//Tick one iteration
func (p *Plant) Tick() {
	p.value++
}

//Get produced plant
func (p *Plant) Get() Produce {
	var pro Produce
	pro.Value = p.value
	pro.Kind = p.kind
	p.value = 0
	return pro
}

func (p *Plant) String() string {
	return strconv.Itoa(p.kind)
}
