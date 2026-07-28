package main

import "os"

type Trigger struct {
	Status     Int32
	CheckTimer Int64
	DelayTimer Int64
	RearmTimer Int64
}

func (trigger *Trigger) Read(file *os.File) {
	trigger.Status.Read(file)
	trigger.CheckTimer.Read(file)
	trigger.DelayTimer.Read(file)
	trigger.RearmTimer.Read(file)
}

func (trigger *Trigger) Append(buf []byte) []byte {
	buf = trigger.Status.Append(buf)
	buf = trigger.CheckTimer.Append(buf)
	buf = trigger.DelayTimer.Append(buf)
	buf = trigger.RearmTimer.Append(buf)
	return buf
}
