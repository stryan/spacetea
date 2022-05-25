package simulator

import (
	"fmt"
	"sort"
	"strconv"
)

//Player is a player controlled mob
type Player struct {
	Resources   map[itemType]int
	Inventory   map[itemType]int
	Craftables  map[itemType]struct{}
	Techs       map[TechID]Tech
	CurrentTile *Tile
	log         []string
	logIndex    int
}

//NewPlayer initializes a player
func NewPlayer() *Player {
	return &Player{Resources: make(map[itemType]int), Techs: make(map[TechID]Tech), Craftables: make(map[itemType]struct{})}
}

func (p *Player) String() string {
	var res string
	res += "Resources: \n"
	var ress []int
	for k, v := range p.Resources {
		if v != 0 {
			//res += fmt.Sprintf("%v: %v\n", GlobalItems[k].Describe(), v)
			ress = append(ress, int(k))
		}
	}
	sort.Ints(ress)
	for _, k := range ress {
		id := itemType(k)
		res += fmt.Sprintf("%v: %v\n", GlobalItems[id], p.Resources[id])
	}
	res += "\nLocation: \n"
	if p.CurrentTile != nil {
		res += p.CurrentTile.String()
	}
	return res
}

func (p *Player) research() {
	for _, tech := range GlobalTechs {
		if _, ok := p.Techs[tech.ID]; ok {
			continue
		}

		i := 0
		for _, v := range tech.Requires {
			req := lookupByName(v.Name)
			if p.Resources[req.ID()] >= v.Value {
				i++
			}
		}
		if i == len(tech.Requires) {
			p.Techs[tech.ID] = tech
			for _, v := range tech.Unlocks {
				itm := lookupByName(v)
				if itm.Type() != emptyObject {
					p.Craftables[itm.ID()] = struct{}{}
				}
			}
			p.Announce(fmt.Sprintf("New Tech: %v", tech.DisplayName))
		}
	}
}

//Announce adds an entry to a players log
func (p *Player) Announce(msg string) {
	p.logIndex++
	p.log = append(p.log, strconv.Itoa(p.logIndex)+" "+msg)
	if len(p.log) > 3 {
		p.log = p.log[1:]
	}
}

//Log returns the player log
func (p *Player) Log() string {
	res := "Log:\n"
	for _, v := range p.log {
		res += v + "\n"
	}
	return res
}
