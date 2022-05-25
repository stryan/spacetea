package simulator

import "github.com/BurntSushi/toml"

//TechID is a tech level
type TechID int

type relation struct {
	name  string
	value int
}

//Tech is a tech level
type Tech struct {
	ID          int        `toml:"techid"`
	DisplayName string     `toml:"display_name"`
	Name        string     `toml:"name"`
	Requires    []relation `toml:"requires"`
	Unlocks     []string   `toml:"unlocks"`
}

const (
	techPulper TechID = iota
)

//GlobalTechList list of all techs
var GlobalTechList = []TechID{techPulper}

//GlobalTechs list of all techs
var GlobalTechs []Tech

type techs struct {
	tech []Tech
}

//LookupTech converts a tech ID to an item ID
func LookupTech(id TechID) ItemEntry {
	switch id {
	case techPulper:
		return converterEntry{convertPulper, 5, "teaConverter", itemPlantTea, itemPlantWood}
	}
	return nil

}

func lookupTechByName(name string) Tech {
	for _, v := range GlobalTechs {
		if v.Name == name {
			return v
		}
	}
	return Tech{}
}

func loadTechs(filename string) {
	var res techs
	_, err := toml.DecodeFile(filename, &res)
	if err != nil {
		panic(err)
	}
	GlobalTechs = res.tech
}
