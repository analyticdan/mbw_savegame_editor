package savegame

import "os"

type PartyRecord struct {
	Valid Int32
	RawId Int32
	Id    Int32
	Party Party
}

func (partyRecord *PartyRecord) Read(file *os.File, gameVersion Int32) {
	partyRecord.Valid.Read(file)
	if partyRecord.Valid == 1 {
		partyRecord.RawId.Read(file)
		partyRecord.Id.Read(file)
		partyRecord.Party.Read(file, gameVersion)
	}
}

func (partyRecord *PartyRecord) Append(buf []byte, gameVersion Int32) []byte {
	buf = partyRecord.Valid.Append(buf)
	if partyRecord.Valid == 1 {
		buf = partyRecord.RawId.Append(buf)
		buf = partyRecord.Id.Append(buf)
		buf = partyRecord.Party.Append(buf, gameVersion)
	}
	return buf
}
