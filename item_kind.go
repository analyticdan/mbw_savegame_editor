package main

import "os"

type ItemKind struct {
	NumSlots Int32
	Slots    []Int64
}

func (itemKind *ItemKind) Read(file *os.File) {
	itemKind.NumSlots.Read(file)
	itemKind.Slots = make([]Int64, itemKind.NumSlots)
	for i := 0; i < len(itemKind.Slots); i++ {
		itemKind.Slots[i].Read(file)
	}
}