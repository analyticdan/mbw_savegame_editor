package main

import "os"

type MapEvent struct {
	_unused0              String
	Type                  Int32
	PositionX             Float
	PositionY             Float
	LandPositionX         Float
	LandPositionY         Float
	_unused1              Float
	_unused2              Float
	AttackerPartyId       Int32
	DefenderPartyId       Int32
	BattleSimulationTimer Int64
	NextBattleSimulation  Float
}

func (mapEvent *MapEvent) Read(file *os.File) {
	mapEvent._unused0.Read(file)
	mapEvent.Type.Read(file)
	mapEvent.PositionX.Read(file)
	mapEvent.PositionY.Read(file)
	mapEvent.LandPositionX.Read(file)
	mapEvent.LandPositionY.Read(file)
	mapEvent._unused1.Read(file)
	mapEvent._unused2.Read(file)
	mapEvent.AttackerPartyId.Read(file)
	mapEvent.DefenderPartyId.Read(file)
	mapEvent.BattleSimulationTimer.Read(file)
	mapEvent.NextBattleSimulation.Read(file)
}

func (mapEvent *MapEvent) Append(buf []byte) ([]byte, error) {
	buf, err := mapEvent._unused0.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapEvent.Type.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapEvent.PositionX.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapEvent.PositionY.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapEvent.LandPositionX.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapEvent.LandPositionY.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapEvent._unused1.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapEvent._unused2.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapEvent.AttackerPartyId.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapEvent.DefenderPartyId.Append(buf)
	if err != nil {
		return buf, err
	}
	buf, err = mapEvent.BattleSimulationTimer.Append(buf)
	if err != nil {
		return buf, err
	}
	return mapEvent.NextBattleSimulation.Append(buf)
}
