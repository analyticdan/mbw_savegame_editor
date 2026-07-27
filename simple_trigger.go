package main

import "os"

type SimpleTrigger struct {
	CheckTimer Int64
}

func (simpleTrigger *SimpleTrigger) Read(file *os.File) {
	simpleTrigger.CheckTimer.Read(file)
}
