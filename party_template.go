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

func (partyTemplate *PartyTemplate) Append(buf []byte) ([]byte, error) {
	buf, err := partyTemplate.NumPartiesCreated.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = partyTemplate.NumPartiesDestroyed.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = partyTemplate.NumPartiesDestroyedByPlayer.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = partyTemplate.NumSlots.Append(buf)
	if err != nil {
		return buf, err
	}
	for i := 0; i < len(partyTemplate.Slots); i++ {
		buf, err = partyTemplate.Slots[i].Append(buf)
		if err != nil {
			return buf, err
		}
	}
	return buf, err
}
