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

func (item *Item) Append(buf []byte) []byte {
	buf = item.ItemKindId.Append(buf)
	buf = item.ItemFlags.Append(buf)
	return buf
}
