package main

import "os"

type PartyTemplate struct {
	NumPartiesCreated           Int32
	NumPartiesDestroyed         Int32
	NumPartiesDestroyedByPlayer Int32
	NumSlots                    Int32
	Slots                       []Int64
}

func (partyTemplate *PartyTemplate) Read(file *os.File) {
	partyTemplate.NumPartiesCreated.Read(file)
	partyTemplate.NumPartiesDestroyed.Read(file)
	partyTemplate.NumPartiesDestroyedByPlayer.Read(file)
	partyTemplate.NumSlots.Read(file)
	partyTemplate.Slots = make([]Int64, partyTemplate.NumSlots)
	for i := 0; i < len(partyTemplate.Slots); i++ {
		partyTemplate.Slots[i].Read(file)
	}
}