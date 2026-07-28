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

func (quest *Quest) Append(buf []byte) ([]byte, error) {
	buf, err := quest.Progression.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = quest.GiverTroopId.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = quest.Number.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = quest.StartDate.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = quest.Title.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = quest.Text.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = quest.Giver.Append(buf)
	if err != nil {
		return buf, err
	}
	for i := 0; i < len(quest.Notes); i++ {
		buf, err = quest.Notes[i].Append(buf)
		if err != nil {
			return buf, err
		}
	}
	buf, err = quest.NumSlots.Append(buf)
	for i := 0; i < len(quest.Slots); i++ {
		buf, err = quest.Slots[i].Append(buf)
		if err != nil {
			return buf, err
		}
	}
	return buf, err
}
