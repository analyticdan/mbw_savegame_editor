package savegame

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

const (
	/*TODO: Maybe move this to its own file with the other tf_flags. */
	heroFlag = UInt64(0x00000010)
	/*TODO: This should be obtained from module.ini. */
	loadRegularTroopInventory = false
)

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
	isHero := troop.Flags&heroFlag != 0
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

func (troop *Troop) Append(buf []byte) []byte {
	buf = troop.NumSlots.Append(buf)
	for i := 0; i < len(troop.Slots); i++ {
		buf = troop.Slots[i].Append(buf)
	}
	for i := 0; i < len(troop.Attributes); i++ {
		buf = troop.Attributes[i].Append(buf)
	}
	for i := 0; i < len(troop.Proficiencies); i++ {
		buf = troop.Proficiencies[i].Append(buf)
	}
	for i := 0; i < len(troop.Skills); i++ {
		buf = troop.Skills[i].Append(buf)
	}
	for i := 0; i < len(troop.Notes); i++ {
		buf = troop.Notes[i].Append(buf)
	}
	buf = troop.Flags.Append(buf)
	buf = troop.SiteIdAndEntryNo.Append(buf)
	buf = troop.SkillPoints.Append(buf)
	buf = troop.AttributePoints.Append(buf)
	buf = troop.ProficiencyPoints.Append(buf)
	buf = troop.Level.Append(buf)
	isHero := troop.Flags&heroFlag != 0
	if isHero || loadRegularTroopInventory {
		buf = troop.Gold.Append(buf)
		buf = troop.Experience.Append(buf)
		buf = troop.Health.Append(buf)
		buf = troop.FactionId.Append(buf)
		for i := 0; i < len(troop.InventoryItems); i++ {
			buf = troop.InventoryItems[i].Append(buf)
		}
		for i := 0; i < len(troop.EquippedItems); i++ {
			buf = troop.EquippedItems[i].Append(buf)
		}
		for i := 0; i < len(troop.FaceKeys); i++ {
			buf = troop.FaceKeys[i].Append(buf)
		}
		buf = troop.Renamed.Append(buf)
		if troop.Renamed {
			buf = troop.Name.Append(buf)
			buf = troop.NamePlural.Append(buf)
		}
	}
	buf = troop.ClassNo.Append(buf)
	return buf
}
