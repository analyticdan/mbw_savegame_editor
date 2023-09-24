package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Bool bool
type Int32 int32
type Int64 int64
type UInt32 uint32
type UInt64 uint64
type Float float32

type String struct {
	NumChars Int32
	Chars    []byte
	Readable string
}

type Header struct {
	MagicNumber   Int32
	GameVersion   Int32
	ModuleVersion Int32
	SavegameName  String
	PlayerName    String
	PlayerLevel   Int32
	Date          Float
}

type Trigger struct {
	Status     Int32
	CheckTimer Int64
	DelayTimer Int64
	RearmTimer Int64
}

type SimpleTrigger struct {
	CheckTimer Int64
}

type Note struct {
	Text              String
	Value             Int32
	TableauMaterialId Int32
	Available         Bool
}

type Quest struct {
	Progression  Int32
	GiverTroopId Int32
	Number       Int32
	StartDate    Float
	Title        String
	Text         String
	Giver        String
	Notes        [16]Note
	NumSlots     Int32
	Slots        []Int64
}

type InfoPage struct {
	Notes [16]Note
}

type Site struct {
	NumSlots Int32
	Slots    []Int64
}

type Faction struct {
	NumSlots  Int32
	Slots     []Int64
	Relations []Float
	Name      String
	Renamed   Bool
	Color     UInt32
	_unused   Int32
	Notes     [16]Note
}

type MapTrack struct {
	PositionX Float
	PositionY Float
	PositionZ Float
	Rotation  Float
	Age       Float
	Flags     Int32
}

type PartyTemplate struct {
	NumPartiesCreated           Int32
	NumPartiesDestroyed         Int32
	NumPartiesDestroyedByPlayer Int32
	NumSlots                    Int32
	Slots                       []Int64
}

type PartyStack struct {
	TroopId          Int32
	NumTroops        Int32
	NumWoundedTroops Int32
	Flags            Int32
}

type Party struct {
	Id                          String
	Name                        String
	Flags                       UInt64
	MenuId                      Int32
	PartyTemplateId             Int32
	FactionId                   Int32
	Personality                 Int32
	DefaultBehavior             Int32
	CurrentBehavior             Int32
	DefaultBehaviorObjectId     Int32
	CurrentBehaviorObjectId     Int32
	InitialPositionX            Float
	InitialPositionY            Float
	TargetPositionX             Float
	TargetPositionY             Float
	PositionX                   Float
	PositionY                   Float
	PositionZ                   Float
	NumStacks                   Int32
	Stacks                      []PartyStack
	Bearing                     Float
	Renamed                     Bool
	ExtraText                   String
	Morale                      Float
	Hunger                      Float
	_unused1                    Float
	PatrolRadius                Float
	Initiative                  Float
	Helpfulness                 Float
	LabelVisible                Int32
	BanditAttraction            Float
	Marshall                    Int32
	IgnorePlayerTimer           Int64
	BannerMapIconId             Int32
	ExtraMapIconId              Int32
	ExtraMapIconUpDownDistance  Float
	ExtraMapIconUpDownFrequency Float
	ExtraMapIconRotateFrequency Float
	ExtraMapIconFadeFrequency   Float
	AttachedToPartyId           Int32
	_unused2                    Int32
	IsAttached                  Bool
	NumAttachedPartyIds         Int32
	AttachedPartyIds            []Int32
	NumParticleSystemIds        Int32
	ParticleSystemIds           []Int32
	Notes                       [16]Note
	NumSlots                    Int32
	Slots                       []Int64
}

type PartyRecord struct {
	Valid Int32
	RawId Int32
	Id    Int32
	Party Party
}

type PlayerPartyStack struct {
	Experience     Float
	NumUpgradeable Int32
	TroopDnas      [32]Int32
}

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

type MapEventRecord struct {
	Valid    Int32
	Id       Int32
	MapEvent MapEvent
}

type Game struct {
	Header                         Header
	GameTime                       UInt64
	RandomSeed                     Int32
	SaveMode                       Int32
	CombatDifficulty               Int32
	CombatDifficultyFriendlies     Int32
	ReduceCombatAi                 Int32
	ReduceCampaignAi               Int32
	CombatSpeed                    Int32
	DateTimer                      Int64
	Hour                           Int32
	Day                            Int32
	Week                           Int32
	Month                          Int32
	Year                           Int32
	_unused0                       Int32
	GlobalCloudAmount              Float
	GlobalHazeAmount               Float
	AverageDifficulty              Float
	AverageDifficultyPeriod        Float
	_unused1                       String
	_unused2                       Bool
	TutorialFlags                  Int32
	DefaultPrisonerPrice           Int32
	EncounteredParty1Id            Int32
	EncounteredParty2Id            Int32
	CurrentMenuId                  Int32
	CurrentSiteId                  Int32
	CurrentEntryNo                 Int32
	CurrentMissionTemplateId       Int32
	PartyCreationMinRandomValue    Int32
	PartyCreationMaxRandomValue    Int32
	GameLog                        String
	_unused3                       [6]Int32
	_unused4                       Int64
	RestPeriod                     Float
	RestTimeSpeed                  Int32
	RestIsInteractive              Int32
	RestRemainAttackable           Int32
	ClassNames                     [9]String
	NumGlobalVariables             Int32
	GlobalVariables                []Int64
	NumTriggers                    Int32
	Triggers                       []Trigger
	NumSimpleTriggers              Int32
	SimpleTriggers                 []SimpleTrigger
	NumQuests                      Int32
	Quests                         []Quest
	NumInfoPages                   Int32
	InfoPages                      []InfoPage
	NumSites                       Int32
	Sites                          []Site
	NumFactions                    Int32
	Factions                       []Faction
	NumMapTracks                   Int32
	MapTracks                      []MapTrack
	NumPartyTemplates              Int32
	PartyTemplates                 []PartyTemplate
	NumPartyRecords                Int32
	NumPartiesCreated              Int32
	PartyRecords                   []PartyRecord
	PlayerPartyStackAdditionalInfo []PlayerPartyStack
	NumMapEventRecords             Int32
	NumMapEventsCreated            Int32
	MapEventRecords                []MapEventRecord
	NumTroops                      Int32
	/*Troops                         []Troop
	_unused5                       [42]Int32
	NumItemKinds                   Int32
	ItemKinds                      []ItemKind
	PlayerFaceKeys0                Int64
	PlayerFaceKeys1                Int64
	PlayerKillCount                Int32
	PlayerWoundedCount             Int32
	PlayerOwnTroopKillCount        Int32
	PlayerOwnTroopWoundedCount     Int32*/
}

func (b *Bool) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, b)
}

func (i *Int32) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, i)
}

func (i *Int64) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, i)
}

func (i *UInt32) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, i)
}

func (i *UInt64) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, i)
}

func (f *Float) Read(file *os.File) {
	binary.Read(file, binary.LittleEndian, f)
}

func (s *String) Read(file *os.File) {
	s.NumChars.Read(file)
	s.Chars = make([]byte, s.NumChars)
	binary.Read(file, binary.LittleEndian, &s.Chars)
	s.Readable = string(s.Chars)
}

func (header *Header) Read(file *os.File) {
	header.MagicNumber.Read(file)
	if header.MagicNumber != 0x52445257 {
		log.Fatal("Magic number not 0x52445257")
	}
	header.GameVersion.Read(file)
	header.ModuleVersion.Read(file)
	header.SavegameName.Read(file)
	header.PlayerName.Read(file)
	header.PlayerLevel.Read(file)
	header.Date.Read(file)
}

func (trigger *Trigger) Read(file *os.File) {
	trigger.Status.Read(file)
	trigger.CheckTimer.Read(file)
	trigger.DelayTimer.Read(file)
	trigger.RearmTimer.Read(file)
}

func (simpleTrigger *SimpleTrigger) Read(file *os.File) {
	simpleTrigger.CheckTimer.Read(file)
}

func (note *Note) Read(file *os.File) {
	note.Text.Read(file)
	note.Value.Read(file)
	note.TableauMaterialId.Read(file)
	note.Available.Read(file)
}

func (quest *Quest) Read(file *os.File) {
	quest.Progression.Read(file)
	quest.GiverTroopId.Read(file)
	quest.Number.Read(file)
	quest.StartDate.Read(file)
	quest.Title.Read(file)
	quest.Text.Read(file)
	quest.Giver.Read(file)
	for i := 0; i < len(quest.Notes); i++ {
		quest.Notes[i].Read(file)
	}
	quest.NumSlots.Read(file)
	quest.Slots = make([]Int64, quest.NumSlots)
	for i := 0; i < len(quest.Slots); i++ {
		quest.Slots[i].Read(file)
	}
}

func (infoPage *InfoPage) Read(file *os.File) {
	for i := 0; i < len(infoPage.Notes); i++ {
		infoPage.Notes[i].Read(file)
	}
}

func (site *Site) Read(file *os.File) {
	site.NumSlots.Read(file)
	site.Slots = make([]Int64, site.NumSlots)
	for i := 0; i < len(site.Slots); i++ {
		site.Slots[i].Read(file)
	}
}

func (faction *Faction) Read(file *os.File) {
	faction.NumSlots.Read(file)
	faction.Slots = make([]Int64, faction.NumSlots)
	for i := 0; i < len(faction.Slots); i++ {
		faction.Slots[i].Read(file)
	}
	for i := 0; i < len(faction.Relations); i++ {
		faction.Relations[i].Read(file)
	}
	faction.Name.Read(file)
	faction.Renamed.Read(file)
	faction.Color.Read(file)
	faction._unused.Read(file)
	for i := 0; i < len(faction.Notes); i++ {
		faction.Notes[i].Read(file)
	}
}

func (mapTrack *MapTrack) Read(file *os.File) {
	mapTrack.PositionX.Read(file)
	mapTrack.PositionY.Read(file)
	mapTrack.PositionZ.Read(file)
	mapTrack.Rotation.Read(file)
	mapTrack.Age.Read(file)
	mapTrack.Flags.Read(file)
}

func (partyTemplate *PartyTemplate) Read(file *os.File) {
	partyTemplate.NumPartiesCreated.Read(file)
	partyTemplate.NumPartiesDestroyed.Read(file)
	partyTemplate.NumPartiesDestroyedByPlayer.Read(file)
	partyTemplate.NumSlots.Read(file)
	partyTemplate.Slots = make([]Int64, partyTemplate.NumSlots)
	for i := 0; i < len(partyTemplate.Slots); i++ {
		partyTemplate.Slots[i].Read(file)
	}
}

func (partyStack *PartyStack) Read(file *os.File) {
	partyStack.TroopId.Read(file)
	partyStack.NumTroops.Read(file)
	partyStack.NumWoundedTroops.Read(file)
	partyStack.Flags.Read(file)
}

func (party *Party) Read(file *os.File, gameVersion Int32) {
	party.Id.Read(file)
	party.Name.Read(file)
	party.Flags.Read(file)
	party.MenuId.Read(file)
	party.PartyTemplateId.Read(file)
	party.FactionId.Read(file)
	party.Personality.Read(file)
	party.DefaultBehavior.Read(file)
	party.CurrentBehavior.Read(file)
	party.DefaultBehaviorObjectId.Read(file)
	party.CurrentBehaviorObjectId.Read(file)
	party.InitialPositionX.Read(file)
	party.InitialPositionY.Read(file)
	party.TargetPositionX.Read(file)
	party.TargetPositionY.Read(file)
	party.PositionX.Read(file)
	party.PositionY.Read(file)
	party.PositionZ.Read(file)

	party.NumStacks.Read(file)
	party.Stacks = make([]PartyStack, party.NumStacks)
	for i := 0; i < len(party.Stacks); i++ {
		party.Stacks[i].Read(file)
	}

	party.Bearing.Read(file)
	party.Renamed.Read(file)
	party.ExtraText.Read(file)
	party.Morale.Read(file)
	party.Hunger.Read(file)
	party._unused1.Read(file)
	party.PatrolRadius.Read(file)
	party.Initiative.Read(file)
	party.Helpfulness.Read(file)
	party.LabelVisible.Read(file)
	party.BanditAttraction.Read(file)
	if (gameVersion >= 900 && gameVersion < 1000) || gameVersion >= 1020 {
		party.Marshall.Read(file)
	}
	party.IgnorePlayerTimer.Read(file)
	party.BannerMapIconId.Read(file)
	if gameVersion >= 1137 {
		party.ExtraMapIconId.Read(file)
		party.ExtraMapIconUpDownDistance.Read(file)
		party.ExtraMapIconUpDownFrequency.Read(file)
		party.ExtraMapIconRotateFrequency.Read(file)
		party.ExtraMapIconFadeFrequency.Read(file)
	}

	party.AttachedToPartyId.Read(file)
	if gameVersion >= 1162 {
		party._unused2.Read(file)
	}
	party.IsAttached.Read(file)
	party.NumAttachedPartyIds.Read(file)
	party.AttachedPartyIds = make([]Int32, party.NumAttachedPartyIds)
	for i := 0; i < len(party.AttachedPartyIds); i++ {
		party.AttachedPartyIds[i].Read(file)
	}

	party.NumParticleSystemIds.Read(file)
	party.ParticleSystemIds = make([]Int32, party.NumParticleSystemIds)
	for i := 0; i < len(party.ParticleSystemIds); i++ {
		party.ParticleSystemIds[i].Read(file)
	}

	for i := 0; i < len(party.Notes); i++ {
		party.Notes[i].Read(file)
	}

	party.NumSlots.Read(file)
	party.Slots = make([]Int64, party.NumSlots)
	for i := 0; i < len(party.Slots); i++ {
		party.Slots[i].Read(file)
	}
}

func (partyRecord *PartyRecord) Read(file *os.File, gameVersion Int32) {
	partyRecord.Valid.Read(file)
	if partyRecord.Valid == 1 {
		partyRecord.RawId.Read(file)
		partyRecord.Id.Read(file)
		partyRecord.Party.Read(file, gameVersion)
	}
}

func (game *Game) Read(file *os.File) {
	game.Header.Read(file)
	game.GameTime.Read(file)
	game.RandomSeed.Read(file)
	game.SaveMode.Read(file)
	if game.Header.GameVersion >= 1137 {
		game.CombatDifficulty.Read(file)
		game.CombatDifficultyFriendlies.Read(file)
		game.ReduceCombatAi.Read(file)
		game.ReduceCampaignAi.Read(file)
		game.CombatSpeed.Read(file)
	}
	game.DateTimer.Read(file)
	game.Hour.Read(file)
	game.Day.Read(file)
	game.Week.Read(file)
	game.Month.Read(file)
	game.Year.Read(file)
	game._unused0.Read(file)
	game.GlobalCloudAmount.Read(file)
	game.GlobalHazeAmount.Read(file)
	game.AverageDifficulty.Read(file)
	game.AverageDifficultyPeriod.Read(file)
	game._unused1.Read(file)
	game._unused2.Read(file)
	game.TutorialFlags.Read(file)
	game.DefaultPrisonerPrice.Read(file)
	game.EncounteredParty1Id.Read(file)
	game.EncounteredParty2Id.Read(file)
	game.CurrentMenuId.Read(file)
	game.CurrentSiteId.Read(file)
	game.CurrentEntryNo.Read(file)
	game.CurrentMissionTemplateId.Read(file)
	game.PartyCreationMinRandomValue.Read(file)
	game.PartyCreationMaxRandomValue.Read(file)
	game.GameLog.Read(file)
	for i := 0; i < len(game._unused3); i++ {
		game._unused3[i].Read(file)
	}
	game._unused4.Read(file)
	game.RestPeriod.Read(file)
	game.RestTimeSpeed.Read(file)
	game.RestIsInteractive.Read(file)
	game.RestRemainAttackable.Read(file)
	for i := 0; i < len(game.ClassNames); i++ {
		game.ClassNames[i].Read(file)
	}

	game.NumGlobalVariables.Read(file)
	game.GlobalVariables = make([]Int64, game.NumGlobalVariables)
	for i := 0; i < len(game.GlobalVariables); i++ {
		game.GlobalVariables[i].Read(file)
	}

	game.NumTriggers.Read(file)
	game.Triggers = make([]Trigger, game.NumTriggers)
	for i := 0; i < len(game.Triggers); i++ {
		game.Triggers[i].Read(file)
	}

	game.NumSimpleTriggers.Read(file)
	game.SimpleTriggers = make([]SimpleTrigger, game.NumSimpleTriggers)
	for i := 0; i < len(game.SimpleTriggers); i++ {
		game.SimpleTriggers[i].Read(file)
	}

	game.NumQuests.Read(file)
	game.Quests = make([]Quest, game.NumQuests)
	for i := 0; i < len(game.Quests); i++ {
		game.Quests[i].Read(file)
	}

	game.NumInfoPages.Read(file)
	game.InfoPages = make([]InfoPage, game.NumInfoPages)
	for i := 0; i < len(game.InfoPages); i++ {
		game.InfoPages[i].Read(file)
	}

	game.NumSites.Read(file)
	game.Sites = make([]Site, game.NumSites)
	for i := 0; i < len(game.Sites); i++ {
		game.Sites[i].Read(file)
	}

	game.NumFactions.Read(file)
	game.Factions = make([]Faction, game.NumFactions)
	for i := 0; i < len(game.Factions); i++ {
		game.Factions[i].Relations = make([]Float, game.NumFactions)
		game.Factions[i].Read(file)
	}

	game.NumMapTracks.Read(file)
	game.MapTracks = make([]MapTrack, game.NumMapTracks)
	for i := 0; i < len(game.MapTracks); i++ {
		game.MapTracks[i].Read(file)
	}

	game.NumPartyTemplates.Read(file)
	game.PartyTemplates = make([]PartyTemplate, game.NumPartyTemplates)
	for i := 0; i < len(game.PartyTemplates); i++ {
		game.PartyTemplates[i].Read(file)
	}

	game.NumPartyRecords.Read(file)
	game.NumPartiesCreated.Read(file)
	game.PartyRecords = make([]PartyRecord, game.NumPartyRecords)
	for i := 0; i < len(game.PartyRecords); i++ {
		game.PartyRecords[i].Read(file, game.Header.GameVersion)
	}
}

func main() {
	path := "/home/daniel/.mbwarband/Savegames/Native/sg00.sav"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	game := Game{}
	game.Read(file)

	bytes, _ := json.MarshalIndent(game.PartyRecords[0], "", "  ")
	fmt.Println(string(bytes))
}
