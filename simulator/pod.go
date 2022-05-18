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
			if v.Maker != nil {
				v.Maker.Tick()
			}
		}
	}
}

//Place an item on a tile
func (p *Pod) Place(item Producer, x, y int) bool {
	if p.Tiles[x][y].Maker == nil {
		p.Tiles[x][y].Maker = item
		return true
	}
	return false
}

//MovePlayer swaps player tiles
func (p *Pod) MovePlayer(x, y, s, t int) bool {
	if oob(x) || oob(y) || oob(s) || oob(t) {
		return false
	}
	if p.Tiles[x][y].User == nil || p.Tiles[s][t].User != nil {
		return false
	}
	p.Tiles[s][t].User = p.Tiles[x][y].User
	p.Tiles[x][y].User = nil
	return true
}

func (p *Pod) String() string {
	var res string
	res += "##########\n"
	for i := range p.Tiles {
		res += "#"
		for _, v := range p.Tiles[i] {
			if v.User != nil {
				res += "@"
			} else if v.Maker != nil {
				res += v.Maker.String()
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
