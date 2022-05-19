package simulator

import (
	"fmt"
)

//Player is a player controlled mob
type Player struct {
	Resources   map[itemType]int
	CurrentTile *Tile
}

//NewPlayer initializes a player
func NewPlayer() *Player {
	return &Player{Resources: make(map[itemType]int)}
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
