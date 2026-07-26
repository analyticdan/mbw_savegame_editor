package main

import "os"

type InfoPage struct {
	Notes [16]Note
}

func (infoPage *InfoPage) Read(file *os.File) {
	for i := 0; i < len(infoPage.Notes); i++ {
		infoPage.Notes[i].Read(file)
	}
}