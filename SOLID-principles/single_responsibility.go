package main

// YOUR PACKAGE HAS A SINGLE RESPONSIBILITY!

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

// in this case a journal.
type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)

	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// delete entry..
}

// seperation of concerns.

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...load from file
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// ...load from url
}

// func main() {
// 	j := Journal{}
// 	j.AddEntry("I had fun today.")
// 	j.AddEntry("I cried today.")
// 	j.AddEntry("I ate today.")
// 	j.AddEntry("I code everyday.")
// 	fmt.Println(j.String())

// 	fmt.Println(entryCount)
// }
