package main

import (
	"log"
	"os"
)

type Header struct {
	MagicNumber   Int32
	GameVersion   Int32
	ModuleVersion Int32
	SavegameName  String
	PlayerName    String
	PlayerLevel   Int32
	Date          Float
}

func (header *Header) Read(file *os.File) {
	header.MagicNumber.Read(file)
	if header.MagicNumber != 0x52445257 {
		log.Fatal("Magic number not 0x52445257")
	}
	header.GameVersion.Read(file)
	header.ModuleVersion.Read(file)
	header.SavegameName.Read(file)
	header.PlayerName.Read(file)
	header.PlayerLevel.Read(file)
	header.Date.Read(file)
}
