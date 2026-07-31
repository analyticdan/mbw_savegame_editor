package savegame

import (
	"os"
)

type InfoPage struct {
	Notes [16]Note
}

func (infoPage *InfoPage) Read(file *os.File) {
	for i := 0; i < len(infoPage.Notes); i++ {
		infoPage.Notes[i].Read(file)
	}
}

func (infoPage *InfoPage) Append(buf []byte) []byte {
	for i := 0; i < len(infoPage.Notes); i++ {
		buf = infoPage.Notes[i].Append(buf)
	}
	return buf
}
