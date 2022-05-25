package simulator

//Pod is a "location" where a player can move
type Pod struct {
	Tiles [8][8]Tile
}

func newPod() *Pod {
	return &Pod{
		Tiles: [8][8]Tile{},
	}
}

//Tick one iteration
func (p *Pod) Tick() {
	for i := range p.Tiles {
		for _, v := range p.Tiles[i] {
			if v.Building != nil {
				if v.Building.Type() == consumerObject {
					obj := v.Building.(*Converter)
					obj.Tick()
				}
				if v.Building.Type() == resourceObject {
					obj := v.Building.(*Resource)
					obj.Tick()
				}
			}
		}
	}
}

//Place an item on a tile
func (p *Pod) Place(item item, x, y int) bool {
	if p.Tiles[x][y].Building == nil {
		p.Tiles[x][y].Building = item
		return true
	}
	return false
}

//Delete removes an item from a tile
func (p *Pod) Delete(x, y int) {
	p.Tiles[x][y].Building = nil
}

//MovePlayer swaps player tiles
func (p *Pod) MovePlayer(x, y, s, t int) *Tile {
	if oob(x) || oob(y) || oob(s) || oob(t) {
		return nil
	}
	if p.Tiles[x][y].User == nil || p.Tiles[s][t].User != nil {
		return nil
	}
	p.Tiles[s][t].User = p.Tiles[x][y].User
	p.Tiles[x][y].User = nil
	return &p.Tiles[s][t]
}

func (p *Pod) String() string {
	var res string
	res += "##########\n"
	for i := range p.Tiles {
		res += "#"
		for _, v := range p.Tiles[i] {
			if v.User != nil {
				res += "@"
			} else if v.Building != nil {
				res += v.Building.Render()
			} else {
				res += "."
			}
		}
		res += "#\n"
	}
	res += "##########"
	return res
}

func oob(i int) bool {
	return (i >= 8 || i < 0)
}
