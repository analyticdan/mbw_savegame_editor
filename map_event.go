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

func (mapEvent MapEvent) Read(file *os.File) {
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
