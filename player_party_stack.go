package main

import "os"

type PlayerPartyStack struct {
	I 			   int
	IsValid        bool
	Experience     Float
	NumUpgradeable Int32
	TroopDnas      [32]Int32
}

func (playerPartyStack *PlayerPartyStack) Read(file *os.File, i int) {
	playerPartyStack.Experience.Read(file)
	playerPartyStack.NumUpgradeable.Read(file)
	if i < 32 {
		for j := 0; j < len(playerPartyStack.TroopDnas); j++ {
			playerPartyStack.TroopDnas[j].Read(file)
		}
	}
}