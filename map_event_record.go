package main

import (
	"os"
)

type MapEventRecord struct {
	Valid    Int32
	Id       Int32
	MapEvent MapEvent
}

func (mapEventRecord *MapEventRecord) Read(file *os.File) {
	mapEventRecord.Valid.Read(file)
	if mapEventRecord.Valid == 1 {
		mapEventRecord.Id.Read(file)
		mapEventRecord.MapEvent.Read(file)
	}
}

func (mapEventRecord *MapEventRecord) Append(buf []byte) ([]byte, error) {
	buf, err := mapEventRecord.Valid.Append(buf)
	if err != nil {
		return buf, err
	}
	if mapEventRecord.Valid == 1 {
		buf, err = mapEventRecord.Id.Append(buf)
		if err != nil {
			return buf, err
		}
		buf, err = mapEventRecord.MapEvent.Append(buf)
		if err != nil {
			return buf, err
		}
	}
	return buf, err
}
