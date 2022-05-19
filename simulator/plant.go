package simulator

import (
	"strconv"
)

//Plant is a plant that grows per tick
type Plant struct {
	kind   int
	value  int
	growth int
}

func newPlant(k int) *Plant {
	return &Plant{
		kind:   k,
		value:  0,
		growth: 0,
	}
}

//Tick one iteration
func (p *Plant) Tick() {
	p.growth++
	if p.growth > 10 {
		p.value++
		p.growth = 0
	}
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
