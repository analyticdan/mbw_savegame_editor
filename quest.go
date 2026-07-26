package main

import "os"

type Quest struct {
	Progression  Int32
	GiverTroopId Int32
	Number       Int32
	StartDate    Float
	Title        String
	Text         String
	Giver        String
	Notes        [16]Note
	NumSlots     Int32
	Slots        []Int64
}

func (quest *Quest) Read(file *os.File) {
	quest.Progression.Read(file)
	quest.GiverTroopId.Read(file)
	quest.Number.Read(file)
	quest.StartDate.Read(file)
	quest.Title.Read(file)
	quest.Text.Read(file)
	quest.Giver.Read(file)
	for i := 0; i < len(quest.Notes); i++ {
		quest.Notes[i].Read(file)
	}
	quest.NumSlots.Read(file)
	quest.Slots = make([]Int64, quest.NumSlots)
	for i := 0; i < len(quest.Slots); i++ {
		quest.Slots[i].Read(file)
	}
}