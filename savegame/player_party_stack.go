package savegame

import "os"

type PlayerPartyStack struct {
	Experience     Float
	NumUpgradeable Int32
	TroopDnas      [32]Int32
}

func (playerPartyStack *PlayerPartyStack) Read(file *os.File, stackIndex int) {
	playerPartyStack.Experience.Read(file)
	playerPartyStack.NumUpgradeable.Read(file)
	if stackIndex < 32 {
		for i := 0; i < len(playerPartyStack.TroopDnas); i++ {
			playerPartyStack.TroopDnas[i].Read(file)
		}
	}
}

func (playerPartyStack *PlayerPartyStack) Append(buf []byte, stackIndex int) []byte {
	buf = playerPartyStack.Experience.Append(buf)
	buf = playerPartyStack.NumUpgradeable.Append(buf)
	if stackIndex < 32 {
		for i := 0; i < len(playerPartyStack.TroopDnas); i++ {
			buf = playerPartyStack.TroopDnas[i].Append(buf)
		}
	}
	return buf
}
