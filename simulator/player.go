package simulator

import (
	"fmt"
	"sort"
	"strconv"
)

//Player is a player controlled mob
type Player struct {
	Resources   map[itemType]int
	Craftables  map[itemType]struct{}
	Techs       map[TechID]Tech
	Pages       map[PageID]JournalPage
	CurrentTile *Tile
	log         []string
	logIndex    int
}

//NewPlayer initializes a player
func NewPlayer() *Player {
	p := &Player{
		Resources:  make(map[itemType]int),
		Techs:      make(map[TechID]Tech),
		Craftables: make(map[itemType]struct{}),
		Pages:      make(map[PageID]JournalPage),
	}
	return p
}

//AddItemByName adds the given amount of the item using the item name
func (p *Player) AddItemByName(name string, value int) {
	obj := lookupByName(name)
	if obj.Type() != emptyObject {
		p.AddItem(obj.ID(), value)
	}
}

//AddItem adds the given amount of the item
func (p *Player) AddItem(i itemType, value int) {
	if v, ok := p.Resources[i]; ok {
		p.Resources[i] = v + value
	} else {
		p.Resources[i] = value
	}
}

//DelItem removes the given ammount of the item
func (p *Player) DelItem(i itemType, value int) {
	if v, ok := p.Resources[i]; ok {
		p.Resources[i] = v - value
	}
}

//DelItemByName removes the given ammount of the item using the item name
func (p *Player) DelItemByName(name string, value int) {
	obj := lookupByName(name)
	if obj.Type() != emptyObject {
		p.DelItem(obj.ID(), value)
	}
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
		res += fmt.Sprintf("%v: %v\n", GlobalItems[id].Describe(), p.Resources[id])
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

func (p *Player) journal() {
	for _, page := range GlobalPages {
		if _, ok := p.Pages[page.PageID]; ok {
			continue
		}

		i := 0
		for _, v := range page.Requires {
			req := lookupByName(v.Name)
			if p.Resources[req.ID()] >= v.Value {
				i++
			}
		}
		if i == len(page.Requires) {
			p.Pages[page.PageID] = page
			p.Announce(fmt.Sprintf("New Journal: %v", page.Title))
		}
	}
}

//Announce adds an entry to a players log
func (p *Player) Announce(msg string) {
	p.logIndex++
	p.log = append(p.log, strconv.Itoa(p.logIndex)+" "+msg)
}

//Log returns the player log
func (p *Player) Log() string {
	res := "Log:\n"
	start := p.logIndex - 4
	if start < 0 {
		start = 0
	}
	end := p.logIndex
	for i := start; i < end; i++ {

		v := p.log[i]
		res += v
		if i != end {
			res += "\n"
		}
	}
	return res
}
