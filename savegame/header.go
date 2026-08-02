package savegame

import (
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
		panic("Magic number not 0x52445257")
	}
	header.GameVersion.Read(file)
	header.ModuleVersion.Read(file)
	header.SavegameName.Read(file)
	header.PlayerName.Read(file)
	header.PlayerLevel.Read(file)
	header.Date.Read(file)
}

func (header *Header) Append(buf []byte) []byte {
	buf = header.MagicNumber.Append(buf)
	buf = header.GameVersion.Append(buf)
	buf = header.ModuleVersion.Append(buf)
	buf = header.SavegameName.Append(buf)
	buf = header.PlayerName.Append(buf)
	buf = header.PlayerLevel.Append(buf)
	buf = header.Date.Append(buf)
	return buf
}
