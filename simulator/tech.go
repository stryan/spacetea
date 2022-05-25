package simulator

import "github.com/BurntSushi/toml"

//TechID is a tech level
type TechID int

type relation struct {
	Name  string
	Value int
}

//Tech is a tech level
type Tech struct {
	ID          TechID     `toml:"techid"`
	DisplayName string     `toml:"displayName"`
	Name        string     `toml:"name"`
	Flavour     string     `toml:"flavour"`
	Requires    []relation `toml:"requires"`
	Unlocks     []string   `toml:"unlocks"`
}

const (
	techPulper TechID = iota
)

//GlobalTechs list of all techs
var GlobalTechs []Tech

type techs struct {
	Tech []Tech
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
	GlobalTechs = res.Tech
}
