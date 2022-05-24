package simulator

import (
	"fmt"
	"strings"
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
	pod.Place(newPlant(itemPlantTea), 4, 4)
	pod.Tiles[0][0].User = player
	player.Announce("Game started")
	return &Simulator{pod, player, 0, 0, 0, make(chan bool)}
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
	cmdS := strings.Split(cmd, " ")
	switch cmdS[0] {
	case "get":
		if cur.Building != nil {
			if cur.Building.Type() == producerObject {
				build := cur.Building.(Producer)
				prod := build.Get()
				if prod.Kind != 0 && prod.Value > 0 {
					s.Player.Resources[prod.Kind] = s.Player.Resources[prod.Kind] + prod.Value
					s.Player.Announce(fmt.Sprintf("Gathered %v %v", prod.Value, Lookup(prod.Kind).Name()))
				}
			}
		}
	case "place":
		if len(cmdS) < 2 {
			return
		}
		item := cmdS[1]
		if item == itemPlantTea.String() {
			res := s.Place.Place(newPlant(itemPlantTea), s.Px, s.Py)
			if res {
				s.Player.Resources[itemPlantTea] = s.Player.Resources[itemPlantTea] - 1
			}
		} else if item == convertPulper.String() {
			res := s.Place.Place(newConverter(convertPulper, s.Player), s.Px, s.Py)
			if res {
				s.Player.Resources[convertPulper] = s.Player.Resources[convertPulper] - 1
			}
		}
	case "craft":
		if len(cmdS) < 2 {
			return
		}
		item := cmdS[1]
		if item == convertPulper.String() {
			if _, ok := s.Player.Techs[techPulper]; ok {
				if s.Player.Resources[itemPlantTea] > 5 {

					s.Player.Resources[convertPulper] = s.Player.Resources[convertPulper] + 1
					s.Player.Resources[itemPlantTea] = s.Player.Resources[itemPlantTea] - 5
				}
			}
		}
	case "left":
		res := s.Place.MovePlayer(s.Px, s.Py, s.Px, s.Py-1)
		if res != nil {
			s.Py = s.Py - 1
			s.Player.CurrentTile = res
		}
	case "right":
		res := s.Place.MovePlayer(s.Px, s.Py, s.Px, s.Py+1)
		if res != nil {
			s.Py = s.Py + 1
			s.Player.CurrentTile = res
		}
	case "up":
		res := s.Place.MovePlayer(s.Px, s.Py, s.Px-1, s.Py)
		if res != nil {
			s.Px = s.Px - 1
			s.Player.CurrentTile = res
		}
	case "down":
		res := s.Place.MovePlayer(s.Px, s.Py, s.Px+1, s.Py)
		if res != nil {
			s.Px = s.Px + 1
			s.Player.CurrentTile = res
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
			s.Player.research()
		}
	}
}
