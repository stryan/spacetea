package simulator

import (
	"fmt"
)

//Tile is a tile
type Tile struct {
	Building Object
	User     *Player
}

func (t *Tile) String() string {
	var res string
	if t.Building != nil {
		res += fmt.Sprintf("There is a %v here\n", t.Building.Describe())
	} else {
		res += "Nothing here"
	}
	return res
}
