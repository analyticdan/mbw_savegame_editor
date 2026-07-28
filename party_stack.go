package main

import "os"

type PartyStack struct {
	TroopId          Int32
	NumTroops        Int32
	NumWoundedTroops Int32
	Flags            Int32
}

func (partyStack *PartyStack) Read(file *os.File) {
	partyStack.TroopId.Read(file)
	partyStack.NumTroops.Read(file)
	partyStack.NumWoundedTroops.Read(file)
	partyStack.Flags.Read(file)
}

func (partyStack *PartyStack) Append(buf []byte) ([]byte, error) {
	buf, err := partyStack.TroopId.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = partyStack.NumTroops.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = partyStack.NumWoundedTroops.Append(buf)
	if err != nil {
		return buf, err
	}
	return partyStack.Flags.Append(buf)
}
