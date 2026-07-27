package main

import "os"

type Troop struct {
	NumSlots          Int32
	Slots             []Int64
	Attributes        [4]Int32
	Proficiencies     [7]Float
	Skills            [6]UInt32
	Notes             [16]Note
	Flags             UInt64
	SiteIdAndEntryNo  Int32
	SkillPoints       Int32
	AttributePoints   Int32
	ProficiencyPoints Int32
	Level             Int32
	Gold              UInt32
	Experience        Int32
	Health            Float
	FactionId         Int32
	InventoryItems    [96]Item
	EquippedItems     [10]Item
	FaceKeys          [4]UInt64
	Renamed           Bool
	Name              String
	NamePlural        String
	ClassNo           Int32
}

func (troop *Troop) Read(file *os.File) {
	troop.NumSlots.Read(file)
	troop.Slots = make([]Int64, troop.NumSlots)
	for i := 0; i < len(troop.Slots); i++ {
		troop.Slots[i].Read(file)
	}
	for i := 0; i < len(troop.Attributes); i++ {
		troop.Attributes[i].Read(file)
	}
	for i := 0; i < len(troop.Proficiencies); i++ {
		troop.Proficiencies[i].Read(file)
	}
	for i := 0; i < len(troop.Skills); i++ {
		troop.Skills[i].Read(file)
	}
	for i := 0; i < len(troop.Notes); i++ {
		troop.Notes[i].Read(file)
	}
	troop.Flags.Read(file)
	troop.SiteIdAndEntryNo.Read(file)
	troop.SkillPoints.Read(file)
	troop.AttributePoints.Read(file)
	troop.ProficiencyPoints.Read(file)
	troop.Level.Read(file)

	/*TODO: Clean up this hardcoded variable.*/
	heroFlag := UInt64(0x00000010)
	isHero := troop.Flags&heroFlag != 0
	/*TODO: This variable should be obtained from module.ini. */
	loadRegularTroopInventory := false
	if isHero || loadRegularTroopInventory {
		troop.Gold.Read(file)
		troop.Experience.Read(file)
		troop.Health.Read(file)
		troop.FactionId.Read(file)
		for i := 0; i < len(troop.InventoryItems); i++ {
			troop.InventoryItems[i].Read(file)
		}
		for i := 0; i < len(troop.EquippedItems); i++ {
			troop.EquippedItems[i].Read(file)
		}
		for i := 0; i < len(troop.FaceKeys); i++ {
			troop.FaceKeys[i].Read(file)
		}
		troop.Renamed.Read(file)
		if troop.Renamed {
			troop.Name.Read(file)
			troop.NamePlural.Read(file)
		}
	}
	troop.ClassNo.Read(file)
}
