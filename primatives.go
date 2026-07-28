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
	err := binary.Read(file, binary.LittleEndian, b)
	if err != nil {
		panic(err)
	}
}

func (b *Bool) Append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, b)
	if err != nil {
		panic(err)
	}
	return buf
}

func (i *Int32) Read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
}

func (i *Int32) Append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
	return buf
}

func (i *Int64) Read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
}

func (i *Int64) Append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
	return buf
}

func (i *UInt32) Read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
}

func (i *UInt32) Append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
	return buf
}

func (i *UInt64) Read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
}

func (i *UInt64) Append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
	return buf
}

func (f *Float) Read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, f)
	if err != nil {
		panic(err)
	}
	/* TODO: Remove this. */
	if JsonDebug && math.IsNaN(float64(*f)) {
		*f = 0
	}
}

func (f *Float) Append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, f)
	if err != nil {
		panic(err)
	}
	return buf
}
