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

func (faction *Faction) Append(buf []byte) []byte {
	buf = faction.NumSlots.Append(buf)
	for i := 0; i < len(faction.Slots); i++ {
		buf = faction.Slots[i].Append(buf)
	}
	for i := 0; i < len(faction.Relations); i++ {
		buf = faction.Relations[i].Append(buf)
	}
	buf = faction.Name.Append(buf)
	buf = faction.Renamed.Append(buf)
	buf = faction.Color.Append(buf)
	buf = faction._unused.Append(buf)
	for i := 0; i < len(faction.Notes); i++ {
		buf = faction.Notes[i].Append(buf)
	}
	return buf
}
