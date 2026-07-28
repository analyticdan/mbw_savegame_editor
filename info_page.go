package main

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

func (infoPage *InfoPage) Append(buf []byte) ([]byte, error) {
	var err error // Need this declaration here to prevent buf being scoped to for-loop.
	for i := 0; i < len(infoPage.Notes); i++ {
		buf, err = infoPage.Notes[i].Append(buf)
		if err != nil {
			return buf, err
		}
	}
	return buf, err
}
