package main

import "os"

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