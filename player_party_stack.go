package main

import "os"

type PlayerPartyStack struct {
	I              int
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
