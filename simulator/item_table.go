package simulator

import "strconv"

//ItemEntry is a human/ui friendly item description
type ItemEntry interface {
	Name() string
	Render() string
	ID() string
}

type itemType int

func (i itemType) String() string {
	return strconv.Itoa(int(i))
}

const (
	itemPlantTea itemType = iota + 1
	itemPlantWood

	convertPulper
)

//GlobalItemList of all items
var GlobalItemList = []itemType{itemPlantTea, itemPlantWood, convertPulper}

//Lookup returns a human friendly item entry
func Lookup(id itemType) ItemEntry {
	switch id {
	case itemPlantTea:
		return plantEntry{itemPlantTea, 1, "tea"}
	case itemPlantWood:
		return plantEntry{itemPlantWood, 10, "wood"}
	case convertPulper:
		return converterEntry{convertPulper, 5, "teaConverter", itemPlantTea, itemPlantWood}
	}
	return nil
}
