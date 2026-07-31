package savegame

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

func (quest *Quest) Append(buf []byte) []byte {
	buf = quest.Progression.Append(buf)
	buf = quest.GiverTroopId.Append(buf)
	buf = quest.Number.Append(buf)
	buf = quest.StartDate.Append(buf)
	buf = quest.Title.Append(buf)
	buf = quest.Text.Append(buf)
	buf = quest.Giver.Append(buf)
	for i := 0; i < len(quest.Notes); i++ {
		buf = quest.Notes[i].Append(buf)
	}
	buf = quest.NumSlots.Append(buf)
	for i := 0; i < len(quest.Slots); i++ {
		buf = quest.Slots[i].Append(buf)
	}
	return buf
}
