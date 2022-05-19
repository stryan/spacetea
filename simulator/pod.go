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
				v.Building.Tick()
			}
		}
	}
}

//Place an item on a tile
func (p *Pod) Place(item Object, x, y int) bool {
	if p.Tiles[x][y].Building == nil {
		p.Tiles[x][y].Building = item
		return true
	}
	return false
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
				res += v.Building.String()
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
