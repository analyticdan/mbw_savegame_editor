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

func (item *Item) Append(buf []byte) ([]byte, error) {
	buf, err := item.ItemKindId.Append(buf)
	if err != nil {
		return buf, err
	}
	return item.ItemFlags.Append(buf)
}
