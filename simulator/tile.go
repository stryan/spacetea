package simulator

import (
	"fmt"
)

//Tile is a tile
type Tile struct {
	Maker Producer
	User  *Player
}

func (t *Tile) String() string {
	var res string
	if t.Maker != nil {
		res += fmt.Sprintf("There is a %v here\n", t.Maker.Describe())
	} else {
		res += "Nothing here"
	}
	return res
}
