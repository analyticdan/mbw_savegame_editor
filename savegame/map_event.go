package savegame

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

func (mapEvent *MapEvent) Append(buf []byte) []byte {
	buf = mapEvent._unused0.Append(buf)
	buf = mapEvent.Type.Append(buf)
	buf = mapEvent.PositionX.Append(buf)
	buf = mapEvent.PositionY.Append(buf)
	buf = mapEvent.LandPositionX.Append(buf)
	buf = mapEvent.LandPositionY.Append(buf)
	buf = mapEvent._unused1.Append(buf)
	buf = mapEvent._unused2.Append(buf)
	buf = mapEvent.AttackerPartyId.Append(buf)
	buf = mapEvent.DefenderPartyId.Append(buf)
	buf = mapEvent.BattleSimulationTimer.Append(buf)
	buf = mapEvent.NextBattleSimulation.Append(buf)
	return buf
}
