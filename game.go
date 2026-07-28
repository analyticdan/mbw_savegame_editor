package main

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

	game.PlayerPartyStackAdditionalInfo = make([]PlayerPartyStack, len(game.PartyRecords[0].Party.Stacks))
	for i := 0; i < len(game.PartyRecords[0].Party.Stacks); i++ {
		/*TODO: Read in troop.txt module data to get true value of isHero.
		Below check uses a hard-coded value based on Native.*/
		troopId := game.PartyRecords[0].Party.Stacks[i].TroopId
		isHero := troopId == 0 || (troopId >= 194 && troopId < 463)

		if isHero {
			game.PlayerPartyStackAdditionalInfo[i].IsValid = false
		} else {
			game.PlayerPartyStackAdditionalInfo[i].IsValid = true
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

func (game *Game) Write() (buf []byte, err error) {
	buf = nil
	buf, err = game.Header.Append(buf)
	buf, err = game.GameTime.Append(buf)
	buf, err = game.RandomSeed.Append(buf)
	buf, err = game.SaveMode.Append(buf)
	if game.Header.GameVersion >= 1137 {
		buf, err = game.CombatDifficulty.Append(buf)
		buf, err = game.CombatDifficultyFriendlies.Append(buf)
		buf, err = game.ReduceCombatAi.Append(buf)
		buf, err = game.ReduceCampaignAi.Append(buf)
		buf, err = game.CombatSpeed.Append(buf)
	}
	buf, err = game.DateTimer.Append(buf)
	buf, err = game.Hour.Append(buf)
	buf, err = game.Day.Append(buf)
	buf, err = game.Week.Append(buf)
	buf, err = game.Month.Append(buf)
	buf, err = game.Year.Append(buf)
	buf, err = game._unused0.Append(buf)
	buf, err = game.GlobalCloudAmount.Append(buf)
	buf, err = game.GlobalHazeAmount.Append(buf)
	buf, err = game.AverageDifficulty.Append(buf)
	buf, err = game.AverageDifficultyPeriod.Append(buf)
	buf, err = game._unused1.Append(buf)
	buf, err = game._unused2.Append(buf)
	buf, err = game.TutorialFlags.Append(buf)
	buf, err = game.DefaultPrisonerPrice.Append(buf)
	buf, err = game.EncounteredParty1Id.Append(buf)
	buf, err = game.EncounteredParty2Id.Append(buf)
	buf, err = game.CurrentMenuId.Append(buf)
	buf, err = game.CurrentSiteId.Append(buf)
	buf, err = game.CurrentEntryNo.Append(buf)
	buf, err = game.CurrentMissionTemplateId.Append(buf)
	buf, err = game.PartyCreationMinRandomValue.Append(buf)
	buf, err = game.PartyCreationMaxRandomValue.Append(buf)
	buf, err = game.GameLog.Append(buf)
	for i := range len(game._unused3) {
		buf, err = game._unused3[i].Append(buf)
	}
	buf, err = game._unused4.Append(buf)
	buf, err = game.RestPeriod.Append(buf)
	buf, err = game.RestTimeSpeed.Append(buf)
	buf, err = game.RestIsInteractive.Append(buf)
	buf, err = game.RestRemainAttackable.Append(buf)
	for i := range len(game.ClassNames) {
		buf, err = game.ClassNames[i].Append(buf)
	}

	buf, err = game.NumGlobalVariables.Append(buf)
	for i := range len(game.GlobalVariables) {
		buf, err = game.GlobalVariables[i].Append(buf)
	}

	buf, err = game.NumTriggers.Append(buf)
	for i := 0; i < len(game.Triggers); i++ {
		buf, err = game.Triggers[i].Append(buf)
	}

	buf, err = game.NumSimpleTriggers.Append(buf)
	for i := 0; i < len(game.SimpleTriggers); i++ {
		buf, err = game.SimpleTriggers[i].Append(buf)
	}

	buf, err = game.NumQuests.Append(buf)
	for i := 0; i < len(game.Quests); i++ {
		buf, err = game.Quests[i].Append(buf)
	}

	buf, err = game.NumInfoPages.Append(buf)
	for i := 0; i < len(game.InfoPages); i++ {
		buf, err = game.InfoPages[i].Append(buf)
	}

	buf, err = game.NumSites.Append(buf)
	for i := 0; i < len(game.Sites); i++ {
		buf, err = game.Sites[i].Append(buf)
	}

	buf, err = game.NumFactions.Append(buf)
	for i := 0; i < len(game.Factions); i++ {
		buf, err = game.Factions[i].Append(buf)
	}

	buf, err = game.NumMapTracks.Append(buf)
	for i := 0; i < len(game.MapTracks); i++ {
		buf, err = game.MapTracks[i].Append(buf)
	}

	buf, err = game.NumPartyTemplates.Append(buf)
	for i := 0; i < len(game.PartyTemplates); i++ {
		buf, err = game.PartyTemplates[i].Append(buf)
	}

	buf, err = game.NumPartyRecords.Append(buf)
	buf, err = game.NumPartiesCreated.Append(buf)
	for i := 0; i < len(game.PartyRecords); i++ {
		buf, err = game.PartyRecords[i].Append(buf, game.Header.GameVersion)
	}

	for i := 0; i < len(game.PartyRecords[0].Party.Stacks); i++ {
		//TODO: Read in troop.txt module data to get true value of isHero.
		//Below check uses a hard-coded value based on Native.
		troopId := game.PartyRecords[0].Party.Stacks[i].TroopId
		isHero := troopId == 0 || (troopId >= 194 && troopId < 463)

		if !isHero {
			buf, err = game.PlayerPartyStackAdditionalInfo[i].Append(buf, i)
		}
	}

	buf, err = game.NumMapEventRecords.Append(buf)
	buf, err = game.NumMapEventsCreated.Append(buf)
	for i := 0; i < len(game.MapEventRecords); i++ {
		buf, err = game.MapEventRecords[i].Append(buf)
	}

	buf, err = game.NumTroops.Append(buf)
	for i := 0; i < len(game.Troops); i++ {
		buf, err = game.Troops[i].Append(buf)
	}

	for i := 0; i < len(game._unused5); i++ {
		buf, err = game._unused5[i].Append(buf)
	}

	buf, err = game.NumItemKinds.Append(buf)
	for i := 0; i < len(game.ItemKinds); i++ {
		buf, err = game.ItemKinds[i].Append(buf)
	}

	buf, err = game.PlayerFaceKeys0.Append(buf)
	buf, err = game.PlayerFaceKeys1.Append(buf)
	buf, err = game.PlayerKillCount.Append(buf)
	buf, err = game.PlayerWoundedCount.Append(buf)
	buf, err = game.PlayerOwnTroopKillCount.Append(buf)
	buf, err = game.PlayerOwnTroopWoundedCount.Append(buf)
	return
}
