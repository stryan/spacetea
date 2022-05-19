package simulator

import (
	"fmt"
	"strconv"
)

//Plant is a plant that grows per tick
type Plant struct {
	kind   itemType
	value  int
	growth int
	rate   int
}

func getPlant(k itemType) plantEntry {
	return Lookup(k).(plantEntry)
}

func newPlant(k itemType) *Plant {
	return &Plant{
		kind:   k,
		value:  0,
		growth: 0,
		rate:   getPlant(k).rate,
	}
}

//Tick one iteration
func (p *Plant) Tick() {
	p.growth++
	if p.growth > p.rate {
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
	return Lookup(p.kind).Render()
}

//Describe returns a human useful string
func (p *Plant) Describe() string {
	return fmt.Sprintf("a %v plant with %v value", Lookup(p.kind).Name(), strconv.Itoa(p.value))
}

//Type returns producer
func (p *Plant) Type() ObjectType {
	return producerObject
}

type plantEntry struct {
	id   itemType
	rate int
	name string
}

func (p plantEntry) Render() string {
	return "w"
}

func (p plantEntry) Name() string {
	return p.name
}

func (p plantEntry) ID() string {
	return p.id.String()
}
