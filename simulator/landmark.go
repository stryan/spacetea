package simulator

import (
	"strconv"

	"github.com/BurntSushi/toml"
)

//LandmarkID is a landmark
type LandmarkID int

//Landmark is a flavour location
type Landmark struct {
	LandmarkID LandmarkID   `toml:"landmarkid"`
	Name       string       `toml:"title"`
	Content    string       `toml:"content"`
	Links      []LandmarkID `toml:"links"`
}

func (l Landmark) ID() string {
	return strconv.Itoa(int(l.LandmarkID))
}

//GlobalPages is a list of all pages
var GlobalLandmarks []Landmark

type landmarks struct {
	Landmark []Landmark
}

func loadLandmarks(filename string) {
	var res landmarks
	_, err := toml.DecodeFile(filename, &res)
	if err != nil {
		panic(err)
	}
	GlobalLandmarks = res.Landmark
}
