package simulator

import (
	"github.com/BurntSushi/toml"
)

//Converter is a object that converts one item into another per tick
type Converter struct {
	Id          itemType `toml:"itemid"`
	Name        string   `toml:"name"`
	DisplayName string   `toml:"displayName"`
	Icon        string   `toml:"icon"`
	Rate        int      `toml:"rate"`
	source      itemType
	SourceName  string `toml:"source"`
	output      itemType
	OutputName  string `toml:"output"`
	owner       *Player
	Costs       []CraftCost
}

type converters struct {
	Converter []Converter
}

//ID returns id
func (c Converter) ID() itemType {
	return c.Id
}

func newConverter(k itemType, o *Player) *Converter {
	var res Converter
	if template, ok := GlobalItems[k]; ok {
		temp := template.(Converter)
		res.DisplayName = temp.DisplayName
		res.Name = temp.Name
		res.Icon = temp.Icon
		res.source = lookupByName(temp.SourceName).ID()
		res.output = lookupByName(temp.OutputName).ID()
		res.Rate = temp.Rate
		res.Costs = temp.Costs
		res.owner = o
		return &res
	}
	return &Converter{}
}

//Tick one iteration
func (c *Converter) Tick() {
	if c.source == 0 {
		c.owner.Resources[c.output] = c.owner.Resources[c.output] + 1
	} else if c.owner.Resources[c.source] > c.Rate {
		c.owner.Resources[c.source] = c.owner.Resources[c.source] - c.Rate
		c.owner.Resources[c.output] = c.owner.Resources[c.output] + 1
	}
}

func (c Converter) String() string {
	return c.Name
}

func (c Converter) Render() string {
	return c.Icon
}

//Describe returns human useful string
func (c Converter) Describe() string {
	return c.DisplayName
}

//Type returns consumer
func (c Converter) Type() ObjectType {
	return consumerObject
}

func loadConverters(filename string) {
	var res converters
	_, err := toml.DecodeFile(filename, &res)
	if err != nil {
		panic(err)
	}
	for _, v := range res.Converter {
		newItem(v.ID(), v)
	}
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
