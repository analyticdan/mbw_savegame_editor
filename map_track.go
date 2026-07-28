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

func (mapTrack *MapTrack) Append(buf []byte) ([]byte, error) {
	buf, err := mapTrack.PositionX.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapTrack.PositionY.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapTrack.PositionZ.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapTrack.Rotation.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapTrack.Age.Append(buf)
	if err != nil {
		return buf, err
	}
	return mapTrack.Flags.Append(buf)
}
