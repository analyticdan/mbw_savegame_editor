package savegame

type Bool bool
type Int32 int32
type Int64 int64
type UInt32 uint32
type UInt64 uint64
type Float float32

type String struct {
	NumChars Int32
	Chars    []byte
}

func (s String) String() string {
	return string(s.Chars)
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
	Unused    Int32
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
	Unused1                     Float
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
	Unused2                     Int32
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
	Unused0               String
	Type                  Int32
	PositionX             Float
	PositionY             Float
	LandPositionX         Float
	LandPositionY         Float
	Unused1               Float
	Unused2               Float
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

type Item struct {
	ItemKindId Int32
	ItemFlags  Int32
}

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

type ItemKind struct {
	NumSlots Int32
	Slots    []Int64
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
	Unused0                        Int32
	GlobalCloudAmount              Float
	GlobalHazeAmount               Float
	AverageDifficulty              Float
	AverageDifficultyPeriod        Float
	Unused1                        String
	Unused2                        Bool
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
	Unused3                        [6]Int32
	Unused4                        Int64
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
	Troops                         []Troop
	Unused5                        [42]Int32
	NumItemKinds                   Int32
	ItemKinds                      []ItemKind
	PlayerFaceKeys0                Int64
	PlayerFaceKeys1                Int64
	PlayerKillCount                Int32
	PlayerWoundedCount             Int32
	PlayerOwnTroopKillCount        Int32
	PlayerOwnTroopWoundedCount     Int32
}
