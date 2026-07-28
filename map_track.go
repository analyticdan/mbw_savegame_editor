package main

import "os"

type MapTrack struct {
	PositionX Float
	PositionY Float
	PositionZ Float
	Rotation  Float
	Age       Float
	Flags     Int32
}

func (mapTrack *MapTrack) Read(file *os.File) {
	mapTrack.PositionX.Read(file)
	mapTrack.PositionY.Read(file)
	mapTrack.PositionZ.Read(file)
	mapTrack.Rotation.Read(file)
	mapTrack.Age.Read(file)
	mapTrack.Flags.Read(file)
}

func (mapTrack *MapTrack) Append(buf []byte) []byte {
	buf = mapTrack.PositionX.Append(buf)
	buf = mapTrack.PositionY.Append(buf)
	buf = mapTrack.PositionZ.Append(buf)
	buf = mapTrack.Rotation.Append(buf)
	buf = mapTrack.Age.Append(buf)
	buf = mapTrack.Flags.Append(buf)
	return buf
}
