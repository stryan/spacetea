package simulator

import (
	"log"

	"github.com/BurntSushi/toml"
)

//Resource is a game resource; can be planted
type Resource struct {
	Id          itemType `toml:"itemid"`
	Name        string   `toml:"name"`
	DisplayName string   `toml:"displayName"`
	Buildable   bool     `toml:"buildable"`
	Rate        int      `toml:"rate"`
	Icon        string   `toml:"icon"`
	value       int
	growth      int
}

type resources struct {
	Resource []Resource
}

func newResource(k itemType) *Resource {
	var res Resource
	if template, ok := GlobalItems[k]; ok {
		temp := template.(Resource)
		res.DisplayName = temp.DisplayName
		res.Id = k
		res.Name = temp.Name
		res.Buildable = temp.Buildable
		res.Rate = temp.Rate
		res.Icon = temp.Icon
		res.value = 0
		res.growth = 0
		return &res
	}
	return &Resource{}
}

func (r *Resource) Tick() {
	if !r.Buildable {
		return
	}
	r.growth++
	if r.growth > r.Rate {
		r.value++
		r.growth = 0
	}
}

func (r *Resource) Get() Produce {
	var pro Produce
	pro.Value = r.value
	pro.Kind = r.Id
	r.value = 0
	return pro
}

func (r Resource) String() string {
	return r.Name
}

func (r Resource) Render() string {
	return r.Icon
}

func (r Resource) Describe() string {
	return r.DisplayName
}
func loadResources(filename string) {
	var res resources
	foo, err := toml.DecodeFile(filename, &res)
	log.Println(foo.Undecoded())
	if err != nil {
		panic(err)
	}
	for _, v := range res.Resource {
		newItem(v.Id, v)
	}
}

func (r Resource) ID() itemType {
	return r.Id
}

func (r Resource) Type() ObjectType {
	return resourceObject
}
