package main

import (
	"encoding/binary"
	"math"
	"os"
)

type Bool bool
type Int32 int32
type Int64 int64
type UInt32 uint32
type UInt64 uint64
type Float float32

func (b *Bool) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, b)
}

func (b *Bool) Append(buf []byte) ([]byte, error) {
	return binary.Append(buf, binary.LittleEndian, b)
}

func (i *Int32) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, i)
}

func (i *Int32) Append(buf []byte) ([]byte, error) {
	return binary.Append(buf, binary.LittleEndian, i)
}

func (i *Int64) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, i)
}

func (i *Int64) Append(buf []byte) ([]byte, error) {
	return binary.Append(buf, binary.LittleEndian, i)
}

func (i *UInt32) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, i)
}

func (i *UInt32) Append(buf []byte) ([]byte, error) {
	return binary.Append(buf, binary.LittleEndian, i)
}

func (i *UInt64) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, i)
}

func (i *UInt64) Append(buf []byte) ([]byte, error) {
	return binary.Append(buf, binary.LittleEndian, i)
}

func (f *Float) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, f)
	if JSON_DEBUG && math.IsNaN(float64(*f)) {
		*f = 0
	}
}

func (f *Float) Append(buf []byte) ([]byte, error) {
	return binary.Append(buf, binary.LittleEndian, f)
}
