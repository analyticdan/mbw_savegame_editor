package main

import "os"

type Faction struct {
	NumSlots  Int32
	Slots     []Int64
	Relations []Float
	Name      String
	Renamed   Bool
	Color     UInt32
	_unused   Int32
	Notes     [16]Note
}

func (faction *Faction) Read(file *os.File) {
	faction.NumSlots.Read(file)
	faction.Slots = make([]Int64, faction.NumSlots)
	for i := 0; i < len(faction.Slots); i++ {
		faction.Slots[i].Read(file)
	}
	for i := 0; i < len(faction.Relations); i++ {
		faction.Relations[i].Read(file)
	}
	faction.Name.Read(file)
	faction.Renamed.Read(file)
	faction.Color.Read(file)
	faction._unused.Read(file)
	for i := 0; i < len(faction.Notes); i++ {
		faction.Notes[i].Read(file)
	}
}