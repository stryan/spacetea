package simulator

import (
	"strconv"

	"github.com/BurntSushi/toml"
)

//PageID is a journal page
type PageID int

//JournalPage is a "flavour event"
type JournalPage struct {
	PageID   PageID     `toml:"pageid"`
	Title    string     `toml:"title"`
	Content  string     `toml:"content"`
	Requires []relation `toml:"requires"`
}

func (j JournalPage) ID() string {
	return strconv.Itoa(int(j.PageID))
}

//GlobalPages is a list of all pages
var GlobalPages []JournalPage

type pages struct {
	JournalPage []JournalPage
}

func loadPages(filename string) {
	var res pages
	_, err := toml.DecodeFile(filename, &res)
	if err != nil {
		panic(err)
	}
	GlobalPages = res.JournalPage
}
