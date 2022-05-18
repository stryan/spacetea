package simulator

import (
	"time"
)

//Simulator contains the main game state
type Simulator struct {
	Place  *Pod
	Player *Player
	Px, Py int
	Time   int
	stop   chan bool
}

//NewSimulator creates a new simulator instance
func NewSimulator() *Simulator {
	pod := newPod()
	player := NewPlayer()
	pod.Place(&Plant{1, 0}, 4, 4)
	pod.Tiles[0][0].User = player
	return &Simulator{pod, NewPlayer(), 0, 0, 0, make(chan bool)}
}

//Start begins the simulation, non blocking
func (s *Simulator) Start() {
	go s.main()
}

//Stop ends the simulation
func (s *Simulator) Stop() {
	s.stop <- true
}

//Input sends a command to the simulator
func (s *Simulator) Input(cmd string) {
	cur := s.Place.Tiles[s.Px][s.Py]
	switch cmd {
	case "get":
		if cur.Maker != nil {
			prod := cur.Maker.Get()
			if prod.Kind != 0 && prod.Value > 0 {
				s.Player.Resources[prod.Kind] = s.Player.Resources[prod.Kind] + prod.Value
			}
		}
	case "left":
		res := s.Place.MovePlayer(s.Px, s.Py, s.Px, s.Py-1)
		if res {
			s.Py = s.Py - 1
		}
	case "right":
		res := s.Place.MovePlayer(s.Px, s.Py, s.Px, s.Py+1)
		if res {
			s.Py = s.Py + 1
		}
	case "up":
		res := s.Place.MovePlayer(s.Px, s.Py, s.Px-1, s.Py)
		if res {
			s.Px = s.Px - 1
		}
	case "down":
		res := s.Place.MovePlayer(s.Px, s.Py, s.Px+1, s.Py)
		if res {
			s.Px = s.Px + 1
		}
	case "quit":
		s.Stop()
	}
}

func (s *Simulator) main() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-s.stop:
			return
		case <-ticker.C:
			s.Time = s.Time + 1
			s.Place.Tick()
		}
	}
}
