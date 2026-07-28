package main

import (
	"os"
)

type Note struct {
	Text              String
	Value             Int32
	TableauMaterialId Int32
	Available         Bool
}

func (note *Note) Read(file *os.File) {
	note.Text.Read(file)
	note.Value.Read(file)
	note.TableauMaterialId.Read(file)
	note.Available.Read(file)
}

func (note *Note) Append(buf []byte) ([]byte, error) {
	buf, err := note.Text.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = note.Value.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = note.TableauMaterialId.Append(buf)
	if err != nil {
		return buf, err
	}
	return note.Available.Append(buf)
}
