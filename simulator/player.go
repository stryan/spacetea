package simulator

import (
	"fmt"
)

//Player is a player controlled mob
type Player struct {
	Resources   map[int]int
	CurrentTile *Tile
}

//NewPlayer initializes a player
func NewPlayer() *Player {
	return &Player{Resources: make(map[int]int)}
}

func (p *Player) String() string {
	var res string
	res += "Resources: \n"
	for i := range p.Resources {
		res += fmt.Sprintf("%v: %v", i, p.Resources[i])
	}
	res += "\nLocation: \n"
	if p.CurrentTile != nil {
		res += p.CurrentTile.String()
	}
	return res
}
