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

func (s *String) Append(buf []byte) ([]byte, error) {
	buf, err := s.NumChars.Append(buf)
	if err != nil {
		return buf, err
	}
	return binary.Append(buf, binary.LittleEndian, &s.Chars)
}
