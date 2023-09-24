package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type String struct {
	NumChars int32
	Chars    []byte
	Readable string
}

type Header struct {
	MagicNumber   int32
	GameVersion   int32
	ModuleVersion int32
	SavegameName  String
	PlayerName    String
	PlayerLevel   int32
	Date          float32
}

type Trigger struct {
	Status     int32
	CheckTimer int64
	DelayTimer int64
	RearmTimer int64
}

type SimpleTrigger struct {
	CheckTimer int64
}

type Note struct {
	Text              String
	Value             int32
	TableauMaterialId int32
	Available         bool
}

type Quest struct {
	Progression  int32
	GiverTroopId int32
	Number       int32
	StartDate    float32
	Title        String
	Text         String
	Giver        String
	Notes        [16]Note
	NumSlots     int32
	Slots        []int64
}

type InfoPage struct {
	Notes [16]Note
}

type Site struct {
	NumSlots int32
	Slots    []int64
}

type Faction struct {
	NumSlots  int32
	Slots     []int64
	Relations []float32
	Name      String
	Renamed   bool
	Color     uint32
	_unused   int32
	Note      [16]Note
}

type MapTrack struct {
	PositionX float32
	PositionY float32
	PositionZ float32
	Rotation  float32
	Age       float32
	Flags     int32
}

type PartyTemplate struct {
	NumPartiesCreated           int32
	NumPartiesDestroyed         int32
	NumPartiesDestroyedByPlayer int32
	NumSlots                    int32
	Slots                       []int64
}

type PartyStack struct {
	TroopId          int32
	NumTroops        int32
	NumWoundedTroops int32
	Flags            int32
}

type Party struct {
	Id                          String
	Name                        String
	Flags                       uint64
	MenuId                      int32
	PartyTemplateId             int32
	FactionId                   int32
	Personality                 int32
	DefaultBehavior             int32
	CurrentBehavior             int32
	DefaultBehaviorObjectId     int32
	CurrentBehaviorObjectId     int32
	InitialPositionX            float32
	InitialPositionY            float32
	TargetPositionX             float32
	TargetPositionY             float32
	PositionX                   float32
	PositionY                   float32
	PositionZ                   float32
	NumStacks                   int32
	Stacks                      []PartyStack
	Bearing                     float32
	Renamed                     bool
	ExtraText                   String
	Morale                      float32
	Hunger                      float32
	_unused1                    float32
	PatrolRadius                float32
	Initiative                  float32
	Helpfulness                 float32
	LabelVisible                int32
	BanditAttraction            float32 // why is this sometimes NaN?
	Marshall                    int32
	IgnorePlayerTimer           int64
	BannerMapIconId             int32
	ExtraMapIconId              int32
	ExtraMapIconUpDownDistance  float32
	ExtraMapIconUpDownFrequency float32
	ExtraMapIconRotateFrequency float32
	ExtraMapIconFadeFrequency   float32
	AttachedToPartyId           int32
	_unused2                    int32
	IsAttached                  bool
	NumAttachedPartyIds         int32
	AttachedPartyIds            []int32
	NumParticleSystemIds        int32
	ParticleSystemIds           []int32
	Notes                       [16]Note
	NumSlots                    int32
	Slots                       []int64
}

type PartyRecord struct {
	Valid int32
	RawId int32
	Id    int32
	Party Party
}

type Save struct {
	Header                      Header
	GameTime                    uint64
	RandomSeed                  int32
	SaveMode                    int32
	CombatDifficulty            int32
	CombatDifficultyFriendlies  int32
	ReduceCombatAi              int32
	ReduceCampaignAi            int32
	CombatSpeed                 int32
	DateTimer                   int64
	Hour                        int32
	Day                         int32
	Week                        int32
	Month                       int32
	Year                        int32
	_unused0                    int32
	GlobalCloudAmount           float32
	GlobalHazeAmount            float32
	AverageDifficulty           float32
	AverageDifficultyPeriod     float32
	_unused1                    String
	_unused2                    bool
	TutorialFlags               int32
	DefaultPrisonerPrice        int32
	EncounteredParty1Id         int32
	EncounteredParty2Id         int32
	CurrentMenuId               int32
	CurrentSiteId               int32
	CurrentEntryNo              int32
	CurrentMissionTemplateId    int32
	PartyCreationMinRandomValue int32
	PartyCreationMaxRandomValue int32
	GameLog                     String
	_unused3                    [6]int32
	_unused4                    int64
	RestPeriod                  float32
	RestTimeSpeed               int32
	RestIsInteractive           int32
	RestRemainAttackable        int32
	ClassNames                  [9]String
	NumGlobalVariables          int32
	GlobalVariables             []int64
	NumTriggers                 int32
	Triggers                    []Trigger
	NumSimpleTriggers           int32
	SimpleTriggers              []SimpleTrigger
	NumQuests                   int32
	Quests                      []Quest
	NumInfoPages                int32
	InfoPages                   []InfoPage
	NumSites                    int32
	Sites                       []Site
	NumFactions                 int32
	Factions                    []Faction
	NumMapTracks                int32
	MapTracks                   []MapTrack
	NumPartyTemplates           int32
	PartyTemplates              []PartyTemplate
	NumPartyRecords             int32
	NumPartiesCreated           int32
	PartyRecords                []PartyRecord
	/*PlayerPartyStackAdditionalInfo []PlayerPartyStack
	NumMapEventRecords             int32
	NumMapEventsCreated            int32
	MapEventRecords                []MapEventRecord
	NumTroops                      int32
	Troops                         []Troop
	_unused5                       [42]int32
	NumItemKinds                   int32
	ItemKinds                      []ItemKind
	PlayerFaceKeys0                int64
	PlayerFaceKeys1                int64
	PlayerKillCount                int32
	PlayerWoundedCount             int32
	PlayerOwnTroopKillCount        int32
	PlayerOwnTroopWoundedCount     int32*/
}

func readString(file *os.File, s *String, not_readable ...bool) {
	binary.Read(file, binary.LittleEndian, &s.NumChars)
	s.Chars = make([]byte, s.NumChars)
	binary.Read(file, binary.LittleEndian, &s.Chars)
	if len(not_readable) == 0 {
		s.Readable = string(s.Chars)
	}
}

func readHeader(file *os.File, header *Header) {
	binary.Read(file, binary.LittleEndian, &header.MagicNumber)
	if header.MagicNumber != 0x52445257 {
		log.Fatal("Magic number not 0x52445257")
	}
	binary.Read(file, binary.LittleEndian, &header.GameVersion)
	binary.Read(file, binary.LittleEndian, &header.ModuleVersion)
	readString(file, &header.SavegameName)
	readString(file, &header.PlayerName)
	binary.Read(file, binary.LittleEndian, &header.PlayerLevel)
	binary.Read(file, binary.LittleEndian, &header.Date)
}

func readNote(file *os.File, note *Note) {
	readString(file, &note.Text)
	binary.Read(file, binary.LittleEndian, &note.Value)
	binary.Read(file, binary.LittleEndian, &note.TableauMaterialId)
	binary.Read(file, binary.LittleEndian, &note.Available)
}

func readQuest(file *os.File, quest *Quest) {
	binary.Read(file, binary.LittleEndian, &quest.Progression)
	binary.Read(file, binary.LittleEndian, &quest.GiverTroopId)
	binary.Read(file, binary.LittleEndian, &quest.Number)
	binary.Read(file, binary.LittleEndian, &quest.StartDate)
	readString(file, &quest.Title)
	readString(file, &quest.Text)
	readString(file, &quest.Giver)
	for i := 0; i < 16; i++ {
		readNote(file, &quest.Notes[i])
	}
	binary.Read(file, binary.LittleEndian, &quest.NumSlots)
	quest.Slots = make([]int64, quest.NumSlots)
	binary.Read(file, binary.LittleEndian, &quest.Slots)

}

func readInfoPage(file *os.File, infoPage *InfoPage) {
	for i := 0; i < 16; i++ {
		readNote(file, &infoPage.Notes[i])
	}
}

func readSite(file *os.File, site *Site) {
	binary.Read(file, binary.LittleEndian, &site.NumSlots)
	site.Slots = make([]int64, site.NumSlots)
	binary.Read(file, binary.LittleEndian, &site.Slots)
}

func readFaction(file *os.File, faction *Faction, numFactions int32) {
	binary.Read(file, binary.LittleEndian, &faction.NumSlots)
	faction.Slots = make([]int64, faction.NumSlots)
	binary.Read(file, binary.LittleEndian, &faction.Slots)
	faction.Relations = make([]float32, numFactions)
	binary.Read(file, binary.LittleEndian, &faction.Relations)
	readString(file, &faction.Name)
	binary.Read(file, binary.LittleEndian, &faction.Renamed)
	binary.Read(file, binary.LittleEndian, &faction.Color)
	binary.Read(file, binary.LittleEndian, &faction._unused)
	for i := 0; i < 16; i++ {
		readNote(file, &faction.Note[i])
	}
}

func readMapTrack(file *os.File, mapTrack *MapTrack) {
	binary.Read(file, binary.LittleEndian, &mapTrack.PositionX)
	binary.Read(file, binary.LittleEndian, &mapTrack.PositionY)
	binary.Read(file, binary.LittleEndian, &mapTrack.PositionZ)
	binary.Read(file, binary.LittleEndian, &mapTrack.Rotation)
	binary.Read(file, binary.LittleEndian, &mapTrack.Age)
	binary.Read(file, binary.LittleEndian, &mapTrack.Flags)
}

func readPartyTemplate(file *os.File, partyTemplate *PartyTemplate) {
	binary.Read(file, binary.LittleEndian, &partyTemplate.NumPartiesCreated)
	binary.Read(file, binary.LittleEndian, &partyTemplate.NumPartiesDestroyed)
	binary.Read(file, binary.LittleEndian, &partyTemplate.NumPartiesDestroyedByPlayer)
	binary.Read(file, binary.LittleEndian, &partyTemplate.NumSlots)
	partyTemplate.Slots = make([]int64, partyTemplate.NumSlots)
	binary.Read(file, binary.LittleEndian, &partyTemplate.Slots)
}

func readPartyRecord(file *os.File, partyRecord *PartyRecord, gameVersion int32) {
	binary.Read(file, binary.LittleEndian, &partyRecord.Valid)
	if partyRecord.Valid == 1 {
		binary.Read(file, binary.LittleEndian, &partyRecord.RawId)
		binary.Read(file, binary.LittleEndian, &partyRecord.Id)
		readParty(file, &partyRecord.Party, gameVersion)
	}
}

func readParty(file *os.File, party *Party, gameVersion int32) {
	readString(file, &party.Id)
	readString(file, &party.Name)

	binary.Read(file, binary.LittleEndian, &party.Flags)
	binary.Read(file, binary.LittleEndian, &party.MenuId)
	binary.Read(file, binary.LittleEndian, &party.PartyTemplateId)
	binary.Read(file, binary.LittleEndian, &party.FactionId)
	binary.Read(file, binary.LittleEndian, &party.Personality)
	binary.Read(file, binary.LittleEndian, &party.DefaultBehavior)
	binary.Read(file, binary.LittleEndian, &party.CurrentBehavior)
	binary.Read(file, binary.LittleEndian, &party.DefaultBehaviorObjectId)
	binary.Read(file, binary.LittleEndian, &party.CurrentBehaviorObjectId)
	binary.Read(file, binary.LittleEndian, &party.InitialPositionX)
	binary.Read(file, binary.LittleEndian, &party.InitialPositionY)
	binary.Read(file, binary.LittleEndian, &party.TargetPositionX)
	binary.Read(file, binary.LittleEndian, &party.TargetPositionY)
	binary.Read(file, binary.LittleEndian, &party.PositionX)
	binary.Read(file, binary.LittleEndian, &party.PositionY)
	binary.Read(file, binary.LittleEndian, &party.PositionZ)
	binary.Read(file, binary.LittleEndian, &party.NumStacks)
	party.Stacks = make([]PartyStack, party.NumStacks)
	for i := 0; i < int(party.NumStacks); i++ {
		readPartyStack(file, &party.Stacks[i])
	}
	binary.Read(file, binary.LittleEndian, &party.Bearing)
	binary.Read(file, binary.LittleEndian, &party.Renamed)
	readString(file, &party.ExtraText)
	binary.Read(file, binary.LittleEndian, &party.Morale)
	binary.Read(file, binary.LittleEndian, &party.Hunger)
	binary.Read(file, binary.LittleEndian, &party._unused1)
	binary.Read(file, binary.LittleEndian, &party.PatrolRadius)
	binary.Read(file, binary.LittleEndian, &party.Initiative)
	binary.Read(file, binary.LittleEndian, &party.Helpfulness)
	binary.Read(file, binary.LittleEndian, &party.LabelVisible)
	binary.Read(file, binary.LittleEndian, &party.BanditAttraction)
	if (gameVersion >= 900 && gameVersion < 1000) || gameVersion >= 1020 {
		binary.Read(file, binary.LittleEndian, &party.Marshall)
	}
	binary.Read(file, binary.LittleEndian, &party.IgnorePlayerTimer)
	binary.Read(file, binary.LittleEndian, &party.BannerMapIconId)
	if gameVersion > 1137 {
		binary.Read(file, binary.LittleEndian, &party.ExtraMapIconId)
		binary.Read(file, binary.LittleEndian, &party.ExtraMapIconUpDownDistance)
		binary.Read(file, binary.LittleEndian, &party.ExtraMapIconUpDownFrequency)
		binary.Read(file, binary.LittleEndian, &party.ExtraMapIconRotateFrequency)
		binary.Read(file, binary.LittleEndian, &party.ExtraMapIconFadeFrequency)
	}
	binary.Read(file, binary.LittleEndian, &party.AttachedToPartyId)
	if gameVersion >= 1162 {
		binary.Read(file, binary.LittleEndian, &party._unused2)
	}

	binary.Read(file, binary.LittleEndian, &party.IsAttached)
	binary.Read(file, binary.LittleEndian, &party.NumAttachedPartyIds)
	party.AttachedPartyIds = make([]int32, party.NumAttachedPartyIds)
	binary.Read(file, binary.LittleEndian, &party.AttachedPartyIds)
	binary.Read(file, binary.LittleEndian, &party.NumParticleSystemIds)
	party.ParticleSystemIds = make([]int32, party.NumParticleSystemIds)
	binary.Read(file, binary.LittleEndian, &party.ParticleSystemIds)
	for i := 0; i < 16; i++ {
		readNote(file, &party.Notes[i])
	}
	binary.Read(file, binary.LittleEndian, &party.NumSlots)
	party.Slots = make([]int64, party.NumSlots)
	binary.Read(file, binary.LittleEndian, &party.Slots)
}

func readPartyStack(file *os.File, partyStack *PartyStack) {
	binary.Read(file, binary.LittleEndian, &partyStack.TroopId)
	binary.Read(file, binary.LittleEndian, &partyStack.NumTroops)
	binary.Read(file, binary.LittleEndian, &partyStack.NumWoundedTroops)
	binary.Read(file, binary.LittleEndian, &partyStack.Flags)
}

func main() {
	path := "/home/daniel/.mbwarband/Savegames/Native/sg01.sav"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	save := Save{}
	readHeader(file, &save.Header)
	binary.Read(file, binary.LittleEndian, &save.GameTime)
	binary.Read(file, binary.LittleEndian, &save.RandomSeed)
	binary.Read(file, binary.LittleEndian, &save.SaveMode)
	if save.Header.GameVersion >= 1137 {
		binary.Read(file, binary.LittleEndian, &save.CombatDifficulty)
		binary.Read(file, binary.LittleEndian, &save.CombatDifficultyFriendlies)
		binary.Read(file, binary.LittleEndian, &save.ReduceCombatAi)
		binary.Read(file, binary.LittleEndian, &save.ReduceCampaignAi)
		binary.Read(file, binary.LittleEndian, &save.CombatSpeed)
	}
	binary.Read(file, binary.LittleEndian, &save.DateTimer)
	binary.Read(file, binary.LittleEndian, &save.Hour)
	binary.Read(file, binary.LittleEndian, &save.Day)
	binary.Read(file, binary.LittleEndian, &save.Week)
	binary.Read(file, binary.LittleEndian, &save.Month)
	binary.Read(file, binary.LittleEndian, &save.Year)
	binary.Read(file, binary.LittleEndian, &save._unused0)
	binary.Read(file, binary.LittleEndian, &save.GlobalCloudAmount)
	binary.Read(file, binary.LittleEndian, &save.GlobalHazeAmount)
	binary.Read(file, binary.LittleEndian, &save.AverageDifficulty)
	binary.Read(file, binary.LittleEndian, &save.AverageDifficultyPeriod)
	readString(file, &save._unused1)
	binary.Read(file, binary.LittleEndian, &save._unused2)
	binary.Read(file, binary.LittleEndian, &save.TutorialFlags)
	binary.Read(file, binary.LittleEndian, &save.DefaultPrisonerPrice)
	binary.Read(file, binary.LittleEndian, &save.EncounteredParty1Id)
	binary.Read(file, binary.LittleEndian, &save.EncounteredParty2Id)
	binary.Read(file, binary.LittleEndian, &save.CurrentMenuId)
	binary.Read(file, binary.LittleEndian, &save.CurrentSiteId)
	binary.Read(file, binary.LittleEndian, &save.CurrentEntryNo)
	binary.Read(file, binary.LittleEndian, &save.CurrentMissionTemplateId)
	binary.Read(file, binary.LittleEndian, &save.PartyCreationMinRandomValue)
	binary.Read(file, binary.LittleEndian, &save.PartyCreationMaxRandomValue)
	readString(file, &save.GameLog)

	binary.Read(file, binary.LittleEndian, &save._unused3)
	binary.Read(file, binary.LittleEndian, &save._unused4)

	binary.Read(file, binary.LittleEndian, &save.RestPeriod)
	binary.Read(file, binary.LittleEndian, &save.RestTimeSpeed)
	binary.Read(file, binary.LittleEndian, &save.RestIsInteractive)
	binary.Read(file, binary.LittleEndian, &save.RestRemainAttackable)

	for i := 0; i < 9; i++ {
		readString(file, &save.ClassNames[i], true)
	}

	binary.Read(file, binary.LittleEndian, &save.NumGlobalVariables)
	save.GlobalVariables = make([]int64, save.NumGlobalVariables)
	binary.Read(file, binary.LittleEndian, &save.GlobalVariables)

	binary.Read(file, binary.LittleEndian, &save.NumTriggers)
	save.Triggers = make([]Trigger, save.NumTriggers)
	binary.Read(file, binary.LittleEndian, &save.Triggers)

	binary.Read(file, binary.LittleEndian, &save.NumSimpleTriggers)
	save.SimpleTriggers = make([]SimpleTrigger, save.NumSimpleTriggers)
	binary.Read(file, binary.LittleEndian, &save.SimpleTriggers)

	binary.Read(file, binary.LittleEndian, &save.NumQuests)
	save.Quests = make([]Quest, save.NumQuests)
	for i := 0; i < int(save.NumQuests); i++ {
		readQuest(file, &save.Quests[i])
	}

	binary.Read(file, binary.LittleEndian, &save.NumInfoPages)
	save.InfoPages = make([]InfoPage, save.NumInfoPages)
	for i := 0; i < int(save.NumInfoPages); i++ {
		readInfoPage(file, &save.InfoPages[i])
	}

	binary.Read(file, binary.LittleEndian, &save.NumSites)
	save.Sites = make([]Site, save.NumSites)
	for i := 0; i < int(save.NumSites); i++ {
		readSite(file, &save.Sites[i])
	}

	binary.Read(file, binary.LittleEndian, &save.NumFactions)
	save.Factions = make([]Faction, save.NumFactions)
	for i := 0; i < int(save.NumFactions); i++ {
		readFaction(file, &save.Factions[i], save.NumFactions)
	}

	binary.Read(file, binary.LittleEndian, &save.NumMapTracks)
	save.MapTracks = make([]MapTrack, save.NumMapTracks)
	for i := 0; i < int(save.NumMapTracks); i++ {
		readMapTrack(file, &save.MapTracks[i])
	}

	binary.Read(file, binary.LittleEndian, &save.NumPartyTemplates)
	save.PartyTemplates = make([]PartyTemplate, save.NumPartyTemplates)
	for i := 0; i < int(save.NumPartyTemplates); i++ {
		readPartyTemplate(file, &save.PartyTemplates[i])
	}

	binary.Read(file, binary.LittleEndian, &save.NumPartyRecords)
	binary.Read(file, binary.LittleEndian, &save.NumPartiesCreated)
	save.PartyRecords = make([]PartyRecord, save.NumPartyRecords)
	for i := 0; i < int(save.NumPartyRecords); i++ {
		readPartyRecord(file, &save.PartyRecords[i], save.Header.GameVersion)
	}

	bytes, _ := json.MarshalIndent(save.PartyRecords, "", "  ")
	fmt.Println(string(bytes))
}
