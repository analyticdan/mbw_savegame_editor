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

func (partyStack *PartyStack) Append(buf []byte) []byte {
	buf = partyStack.TroopId.Append(buf)
	buf = partyStack.NumTroops.Append(buf)
	buf = partyStack.NumWoundedTroops.Append(buf)
	buf = partyStack.Flags.Append(buf)
	return buf
}
