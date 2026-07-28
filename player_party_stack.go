package main

import "os"

type PlayerPartyStack struct {
	IsValid        bool
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

func (playerPartyStack *PlayerPartyStack) Append(buf []byte, stackIndex int) ([]byte, error) {
	buf, err := playerPartyStack.Experience.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = playerPartyStack.NumUpgradeable.Append(buf)
	if err != nil {
		return buf, err
	}
	if stackIndex < 32 {
		for i := 0; i < len(playerPartyStack.TroopDnas); i++ {
			buf, err = playerPartyStack.TroopDnas[i].Append(buf)
			if err != nil {
				return buf, err
			}
		}
	}
	return buf, err
}
