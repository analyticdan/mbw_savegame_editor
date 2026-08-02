package savegame

import (
	"os"
)

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
	Troops                         []Troop
	_unused5                       [42]Int32
	NumItemKinds                   Int32
	ItemKinds                      []ItemKind
	PlayerFaceKeys0                Int64
	PlayerFaceKeys1                Int64
	PlayerKillCount                Int32
	PlayerWoundedCount             Int32
	PlayerOwnTroopKillCount        Int32
	PlayerOwnTroopWoundedCount     Int32
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
	playerParty := game.PartyRecords[0].Party
	game.PlayerPartyStackAdditionalInfo = make([]PlayerPartyStack, len(playerParty.Stacks))
	for i := 0; i < len(playerParty.Stacks); i++ {
		if hasAdditionalInfo(playerParty, i) {
			game.PlayerPartyStackAdditionalInfo[i].Read(file, i)
		}
	}
	game.NumMapEventRecords.Read(file)
	game.NumMapEventsCreated.Read(file)
	game.MapEventRecords = make([]MapEventRecord, game.NumMapEventRecords)
	for i := 0; i < len(game.MapEventRecords); i++ {
		game.MapEventRecords[i].Read(file)
	}
	game.NumTroops.Read(file)
	game.Troops = make([]Troop, game.NumTroops)
	for i := 0; i < len(game.Troops); i++ {
		game.Troops[i].Read(file)
	}
	for i := 0; i < len(game._unused5); i++ {
		game._unused5[i].Read(file)
	}
	game.NumItemKinds.Read(file)
	game.ItemKinds = make([]ItemKind, game.NumItemKinds)
	for i := 0; i < len(game.ItemKinds); i++ {
		game.ItemKinds[i].Read(file)
	}
	game.PlayerFaceKeys0.Read(file)
	game.PlayerFaceKeys1.Read(file)
	game.PlayerKillCount.Read(file)
	game.PlayerWoundedCount.Read(file)
	game.PlayerOwnTroopKillCount.Read(file)
	game.PlayerOwnTroopWoundedCount.Read(file)
}

func (game *Game) Write(buf []byte) []byte {
	buf = game.Header.Append(buf)
	buf = game.GameTime.Append(buf)
	buf = game.RandomSeed.Append(buf)
	buf = game.SaveMode.Append(buf)
	if game.Header.GameVersion >= 1137 {
		buf = game.CombatDifficulty.Append(buf)
		buf = game.CombatDifficultyFriendlies.Append(buf)
		buf = game.ReduceCombatAi.Append(buf)
		buf = game.ReduceCampaignAi.Append(buf)
		buf = game.CombatSpeed.Append(buf)
	}
	buf = game.DateTimer.Append(buf)
	buf = game.Hour.Append(buf)
	buf = game.Day.Append(buf)
	buf = game.Week.Append(buf)
	buf = game.Month.Append(buf)
	buf = game.Year.Append(buf)
	buf = game._unused0.Append(buf)
	buf = game.GlobalCloudAmount.Append(buf)
	buf = game.GlobalHazeAmount.Append(buf)
	buf = game.AverageDifficulty.Append(buf)
	buf = game.AverageDifficultyPeriod.Append(buf)
	buf = game._unused1.Append(buf)
	buf = game._unused2.Append(buf)
	buf = game.TutorialFlags.Append(buf)
	buf = game.DefaultPrisonerPrice.Append(buf)
	buf = game.EncounteredParty1Id.Append(buf)
	buf = game.EncounteredParty2Id.Append(buf)
	buf = game.CurrentMenuId.Append(buf)
	buf = game.CurrentSiteId.Append(buf)
	buf = game.CurrentEntryNo.Append(buf)
	buf = game.CurrentMissionTemplateId.Append(buf)
	buf = game.PartyCreationMinRandomValue.Append(buf)
	buf = game.PartyCreationMaxRandomValue.Append(buf)
	buf = game.GameLog.Append(buf)
	for i := 0; i < len(game._unused3); i++ {
		buf = game._unused3[i].Append(buf)
	}
	buf = game._unused4.Append(buf)
	buf = game.RestPeriod.Append(buf)
	buf = game.RestTimeSpeed.Append(buf)
	buf = game.RestIsInteractive.Append(buf)
	buf = game.RestRemainAttackable.Append(buf)
	for i := 0; i < len(game.ClassNames); i++ {
		buf = game.ClassNames[i].Append(buf)
	}
	buf = game.NumGlobalVariables.Append(buf)
	for i := 0; i < len(game.GlobalVariables); i++ {
		buf = game.GlobalVariables[i].Append(buf)
	}
	buf = game.NumTriggers.Append(buf)
	for i := 0; i < len(game.Triggers); i++ {
		buf = game.Triggers[i].Append(buf)
	}
	buf = game.NumSimpleTriggers.Append(buf)
	for i := 0; i < len(game.SimpleTriggers); i++ {
		buf = game.SimpleTriggers[i].Append(buf)
	}
	buf = game.NumQuests.Append(buf)
	for i := 0; i < len(game.Quests); i++ {
		buf = game.Quests[i].Append(buf)
	}
	buf = game.NumInfoPages.Append(buf)
	for i := 0; i < len(game.InfoPages); i++ {
		buf = game.InfoPages[i].Append(buf)
	}
	buf = game.NumSites.Append(buf)
	for i := 0; i < len(game.Sites); i++ {
		buf = game.Sites[i].Append(buf)
	}
	buf = game.NumFactions.Append(buf)
	for i := 0; i < len(game.Factions); i++ {
		buf = game.Factions[i].Append(buf)
	}
	buf = game.NumMapTracks.Append(buf)
	for i := 0; i < len(game.MapTracks); i++ {
		buf = game.MapTracks[i].Append(buf)
	}
	buf = game.NumPartyTemplates.Append(buf)
	for i := 0; i < len(game.PartyTemplates); i++ {
		buf = game.PartyTemplates[i].Append(buf)
	}
	buf = game.NumPartyRecords.Append(buf)
	buf = game.NumPartiesCreated.Append(buf)
	for i := 0; i < len(game.PartyRecords); i++ {
		buf = game.PartyRecords[i].Append(buf, game.Header.GameVersion)
	}
	playerParty := game.PartyRecords[0].Party
	for i := 0; i < len(game.PlayerPartyStackAdditionalInfo); i++ {
		if hasAdditionalInfo(playerParty, i) {
			buf = game.PlayerPartyStackAdditionalInfo[i].Append(buf, i)
		}
	}
	buf = game.NumMapEventRecords.Append(buf)
	buf = game.NumMapEventsCreated.Append(buf)
	for i := 0; i < len(game.MapEventRecords); i++ {
		buf = game.MapEventRecords[i].Append(buf)
	}
	buf = game.NumTroops.Append(buf)
	for i := 0; i < len(game.Troops); i++ {
		buf = game.Troops[i].Append(buf)
	}
	for i := 0; i < len(game._unused5); i++ {
		buf = game._unused5[i].Append(buf)
	}
	buf = game.NumItemKinds.Append(buf)
	for i := 0; i < len(game.ItemKinds); i++ {
		buf = game.ItemKinds[i].Append(buf)
	}
	buf = game.PlayerFaceKeys0.Append(buf)
	buf = game.PlayerFaceKeys1.Append(buf)
	buf = game.PlayerKillCount.Append(buf)
	buf = game.PlayerWoundedCount.Append(buf)
	buf = game.PlayerOwnTroopKillCount.Append(buf)
	buf = game.PlayerOwnTroopWoundedCount.Append(buf)
	return buf
}

func hasAdditionalInfo(playerParty Party, i int) bool {
	troopId := playerParty.Stacks[i].TroopId
	/*TODO: isHero should be calculated based on troop.txt module data.
	Below uses a hard-coded value based on Native.*/
	isHero := troopId == 0 || (troopId >= 194 && troopId < 463)
	return !isHero
}
