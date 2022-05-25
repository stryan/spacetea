package simulator

import (
	"fmt"
)

//Tile is a tile
type Tile struct {
	Building item
	User     *Player
}

func (t *Tile) String() string {
	var res string
	if t.Building != nil {
		if t.Building.Type() == resourceObject {
			obj := t.Building.(*Resource)
			res += fmt.Sprintf("There is a %v here with value %v\n", obj.Describe(), obj.value)
		} else if t.Building.Type() == consumerObject {
			obj := t.Building.(*Converter)
			res += fmt.Sprintf("There is a %v here\n", obj.Describe())
		}
	} else {
		res += "Nothing here"
	}
	return res
}
