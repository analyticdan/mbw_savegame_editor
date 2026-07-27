package main

import "os"

type Item struct {
	ItemKindId Int32
	ItemFlags  Int32
}

func (item *Item) Read(file *os.File) {
	item.ItemKindId.Read(file)
	item.ItemFlags.Read(file)
}
