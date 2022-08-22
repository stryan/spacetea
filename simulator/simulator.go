package simulator

import (
	"fmt"
	"log"
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

func Demo(s *Simulator) {
	pod := newPod()
	player := NewPlayer()
	pod.Place(newResource(lookupByName("tea").ID()), 4, 4)
	player.AddItem(itemType(1), 30)
	player.AddItem(itemType(3), 5)
	pod.Tiles[0][0].User = player
	player.Announce("Game started")
	s.Place = pod
	s.Player = player
}

//NewSimulator creates a new simulator instance
func NewSimulator() *Simulator {
	log.Println("loading items")
	initItems()
	log.Println("loading techs")
	loadTechs("data/tech.toml")
	log.Println("loading resources")
	loadResources("data/resources.toml")
	log.Println("loading converters")
	loadConverters("data/converters.toml")
	log.Println("loading journal")
	loadPages("data/journal.toml")
	log.Println("loading landmarks")
	loadLandmarks("data/landmark.toml")
	if len(GlobalItems) < 1 {
		panic("Loaded items but nothing in global items table")
	}
	if len(GlobalTechs) < 1 {
		panic("Loaded items but nothing in global items table")
	}
	if len(GlobalPages) < 1 {
		panic("Loaded journal but no pages in table")
	}
	if len(GlobalLandmarks) < 1 {
		panic("Loaded landmarks but nothing in table")
	}

	return &Simulator{nil, nil, 0, 0, 0, make(chan bool)}
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
			if cur.Building.Type() == resourceObject {
				build := cur.Building.(*Resource)
				prod := build.Get()
				if prod.Kind != 0 && prod.Value > 0 {
					s.Player.AddItem(prod.Kind, prod.Value)
					s.Player.Announce(fmt.Sprintf("Gathered %v %v", prod.Value, GlobalItems[prod.Kind].Describe()))
				}
			}
		}
	case "pickup":
		if cur.Building != nil {
			s.Player.AddItem(cur.Building.ID(), 1)
			s.Place.Delete(s.Px, s.Py)
		}
	case "destroy":
		s.Place.Delete(s.Px, s.Py)
	case "place":
		if len(cmdS) < 2 {
			return
		}
		item := cmdS[1]
		s.Player.Announce(fmt.Sprintf("placing %v", item))
		obj := lookupByName(item)
		switch obj.Type() {
		case emptyObject:
			return
		//case producerObject:
		//obj = obj.(Producer)
		case consumerObject:
			obj2 := obj.(Converter)
			res := s.Place.Place(newConverter(obj2.ID(), s.Player), s.Px, s.Py)
			if res {
				s.Player.DelItem(obj2.ID(), 1)
			}
		case resourceObject:
			obj2 := obj.(Resource)
			if obj2.Buildable {
				res := s.Place.Place(newResource(obj2.ID()), s.Px, s.Py)
				if res {
					s.Player.DelItem(obj2.ID(), 1)
				}
			}
		}
	case "craft":
		if len(cmdS) < 2 {
			return
		}
		item := cmdS[1]
		s.Player.Announce(fmt.Sprintf("Crafting %v", item))
		obj := lookupByName(item)
		switch obj.Type() {
		case emptyObject:
			return
		case consumerObject:
			obj2 := obj.(Converter)
			i := 0
			for _, v := range obj2.Costs {
				if s.Player.Resources[lookupByName(v.Name).ID()] >= v.Value {
					i++
				}
			}
			if i == len(obj2.Costs) {
				for _, v := range obj2.Costs {
					s.Player.DelItemByName(v.Name, v.Value)
				}
				s.Player.AddItemByName(obj2.String(), 1)
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
			s.Player.journal()
		}
	}
}
