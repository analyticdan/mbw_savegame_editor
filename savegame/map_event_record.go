package savegame

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

func (mapEventRecord *MapEventRecord) Append(buf []byte) []byte {
	buf = mapEventRecord.Valid.Append(buf)
	if mapEventRecord.Valid == 1 {
		buf = mapEventRecord.Id.Append(buf)
		buf = mapEventRecord.MapEvent.Append(buf)
	}
	return buf
}
