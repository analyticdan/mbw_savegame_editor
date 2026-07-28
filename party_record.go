package main

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

func (partyRecord *PartyRecord) Append(buf []byte, gameVersion Int32) ([]byte, error) {
	buf, err := partyRecord.Valid.Append(buf)
	if err != nil {
		return buf, err
	}
	if partyRecord.Valid == 1 {
		buf, err = partyRecord.RawId.Append(buf)
		if err != nil {
			return buf, err
		}
		buf, err = partyRecord.Id.Append(buf)
		if err != nil {
			return buf, err
		}
		buf, err = partyRecord.Party.Append(buf, gameVersion)
		if err != nil {
			return buf, err
		}
	}
	return buf, err
}
