package main

import (
	"encoding/binary"
	"os"
)

type String struct {
	NumChars Int32
	Chars    []byte
	Readable string
}

func (s *String) Read(file *os.File) {
	s.NumChars.Read(file)
	s.Chars = make([]byte, s.NumChars)
	binary.Read(file, binary.LittleEndian, &s.Chars)
	s.Readable = string(s.Chars)
}
