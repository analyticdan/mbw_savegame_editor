package main

import "os"

type Site struct {
	NumSlots Int32
	Slots    []Int64
}

func (site *Site) Read(file *os.File) {
	site.NumSlots.Read(file)
	site.Slots = make([]Int64, site.NumSlots)
	for i := 0; i < len(site.Slots); i++ {
		site.Slots[i].Read(file)
	}
}