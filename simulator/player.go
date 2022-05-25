package simulator

import (
	"fmt"
	"strconv"
)

//Player is a player controlled mob
type Player struct {
	Resources   map[itemType]int
	Craftables  map[itemType]struct{}
	Techs       map[TechID]struct{}
	CurrentTile *Tile
	log         []string
	logIndex    int
}

//NewPlayer initializes a player
func NewPlayer() *Player {
	return &Player{Resources: make(map[itemType]int), Techs: make(map[TechID]struct{})}
}

func (p *Player) String() string {
	var res string
	res += "Resources: \n"
	for _, i := range GlobalItemList {
		if p.Resources[i] != 0 {
			res += fmt.Sprintf("%v: %v\n", Lookup(i).Name(), p.Resources[i])
		}
	}
	res += "\nLocation: \n"
	if p.CurrentTile != nil {
		res += p.CurrentTile.String()
	}
	return res
}

func (p *Player) research() {
	for k, v := range p.Resources {
		if k == itemPlantTea && v > 10 {
			if _, ok := p.Techs[techPulper]; !ok {
				p.Techs[techPulper] = struct{}{}
				p.Announce("New Tech: Pulper")
			}
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
