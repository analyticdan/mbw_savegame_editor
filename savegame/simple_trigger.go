package savegame

import "os"

type SimpleTrigger struct {
	CheckTimer Int64
}

func (simpleTrigger *SimpleTrigger) Read(file *os.File) {
	simpleTrigger.CheckTimer.Read(file)
}

func (simpleTrigger *SimpleTrigger) Append(buf []byte) []byte {
	buf = simpleTrigger.CheckTimer.Append(buf)
	return buf
}
