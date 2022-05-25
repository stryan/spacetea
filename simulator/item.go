package simulator

import "strconv"

//GlobalItems table
var GlobalItems map[itemType]item
var nameToItem map[string]itemType

type item interface {
	ID() itemType
	Type() ObjectType
	Render() string
	String() string
	Describe() string
}

//ItemEntry is a human/ui friendly item description
type ItemEntry interface {
	String() string
	Render() string
	ID() itemType
}

type itemType int

func (i itemType) String() string {
	return strconv.Itoa(int(i))
}

type empty struct {
}

func (e empty) ID() itemType {
	return itemType(0)
}

func (e empty) Type() ObjectType {
	return emptyObject
}

func (e empty) String() string {
	return ""
}

func (e empty) Render() string {
	return ""
}
func (e empty) Describe() string {
	return "an empty item"
}

func initItems() {
	GlobalItems = make(map[itemType]item)
	nameToItem = make(map[string]itemType)
}

func newItem(id itemType, obj item) {
	if _, ok := GlobalItems[id]; ok {
		panic("trying to add item that already exists")
	}
	if id == 0 || obj.String() == "" {
		panic("trying to add undeclared empty item")
	}
	GlobalItems[id] = obj
	nameToItem[obj.String()] = id
}

func lookupByName(name string) item {
	if res, ok := nameToItem[name]; ok {
		return GlobalItems[res]
	}
	return empty{}
}
