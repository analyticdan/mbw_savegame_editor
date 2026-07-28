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

func (trigger *Trigger) Append(buf []byte) ([]byte, error) {
	buf, err := trigger.Status.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = trigger.CheckTimer.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = trigger.DelayTimer.Append(buf)
	if err != nil {
		return buf, err
	}
	return trigger.RearmTimer.Append(buf)
}
