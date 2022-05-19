package simulator

import (
	"fmt"
)

//Converter is a object that converts one item into another per tick
type Converter struct {
	kind   itemType
	rate   int
	source itemType
	output itemType
	owner  *Player
}

func newConverter(k itemType, o *Player) *Converter {
	return &Converter{
		kind:   k,
		rate:   getConverter(k).rate,
		source: getConverter(k).source,
		output: getConverter(k).output,
		owner:  o,
	}
}

//Tick one iteration
func (c *Converter) Tick() {
	if c.owner.Resources[c.source] > c.rate {
		c.owner.Resources[c.source] = c.owner.Resources[c.source] - c.rate
		c.owner.Resources[c.output] = c.owner.Resources[c.output] + 1
	}
}

func (c *Converter) String() string {
	return Lookup(c.kind).Render()
}

//Describe returns human useful string
func (c *Converter) Describe() string {
	output := getConverter(c.kind).output
	return fmt.Sprintf("a %v converter that outputs %v", Lookup(c.kind).Name(), Lookup(output).Name())
}

//Type returns consumer
func (c *Converter) Type() ObjectType {
	return consumerObject
}

func getConverter(k itemType) converterEntry {
	return Lookup(k).(converterEntry)
}

type converterEntry struct {
	id     itemType
	rate   int
	name   string
	source itemType
	output itemType
}

func (c converterEntry) Name() string {
	return c.name
}

func (c converterEntry) Render() string {
	return "m"
}

func (c converterEntry) ID() string {
	return c.id.String()
}
