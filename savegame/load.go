package savegame

import (
	"encoding/binary"
	"math"
	"os"
)

func (b *Bool) read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, b)
	if err != nil {
		panic(err)
	}
}

func (i *Int32) read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
}

func (i *Int64) read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
}

func (i *UInt32) read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
}

func (i *UInt64) read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
}

func (f *Float) read(file *os.File) {
	err := binary.Read(file, binary.LittleEndian, f)
	if err != nil {
		panic(err)
	}
	if DisableNaN && math.IsNaN(float64(*f)) {
		*f = 0
	}
}

func (s *String) read(file *os.File) {
	s.NumChars.read(file)
	s.Chars = make([]byte, s.NumChars)
	binary.Read(file, binary.LittleEndian, s.Chars)
}

func (header *Header) read(file *os.File) {
	header.MagicNumber.read(file)
	if header.MagicNumber != 0x52445257 {
		panic("Magic number not 0x52445257")
	}
	header.GameVersion.read(file)
	header.ModuleVersion.read(file)
	header.SavegameName.read(file)
	header.PlayerName.read(file)
	header.PlayerLevel.read(file)
	header.Date.read(file)
}

func (trigger *Trigger) read(file *os.File) {
	trigger.Status.read(file)
	trigger.CheckTimer.read(file)
	trigger.DelayTimer.read(file)
	trigger.RearmTimer.read(file)
}

func (simpleTrigger *SimpleTrigger) read(file *os.File) {
	simpleTrigger.CheckTimer.read(file)
}

func (note *Note) read(file *os.File) {
	note.Text.read(file)
	note.Value.read(file)
	note.TableauMaterialId.read(file)
	note.Available.read(file)
}

func (quest *Quest) read(file *os.File) {
	quest.Progression.read(file)
	quest.GiverTroopId.read(file)
	quest.Number.read(file)
	quest.StartDate.read(file)
	quest.Title.read(file)
	quest.Text.read(file)
	quest.Giver.read(file)
	for i := 0; i < len(quest.Notes); i++ {
		quest.Notes[i].read(file)
	}
	quest.NumSlots.read(file)
	quest.Slots = make([]Int64, quest.NumSlots)
	for i := 0; i < len(quest.Slots); i++ {
		quest.Slots[i].read(file)
	}
}

func (infoPage *InfoPage) read(file *os.File) {
	for i := 0; i < len(infoPage.Notes); i++ {
		infoPage.Notes[i].read(file)
	}
}

func (site *Site) read(file *os.File) {
	site.NumSlots.read(file)
	site.Slots = make([]Int64, site.NumSlots)
	for i := 0; i < len(site.Slots); i++ {
		site.Slots[i].read(file)
	}
}

func (faction *Faction) read(file *os.File) {
	faction.NumSlots.read(file)
	faction.Slots = make([]Int64, faction.NumSlots)
	for i := 0; i < len(faction.Slots); i++ {
		faction.Slots[i].read(file)
	}
	for i := 0; i < len(faction.Relations); i++ {
		faction.Relations[i].read(file)
	}
	faction.Name.read(file)
	faction.Renamed.read(file)
	faction.Color.read(file)
	faction.Unused.read(file)
	for i := 0; i < len(faction.Notes); i++ {
		faction.Notes[i].read(file)
	}
}

func (mapTrack *MapTrack) read(file *os.File) {
	mapTrack.PositionX.read(file)
	mapTrack.PositionY.read(file)
	mapTrack.PositionZ.read(file)
	mapTrack.Rotation.read(file)
	mapTrack.Age.read(file)
	mapTrack.Flags.read(file)
}

func (partyTemplate *PartyTemplate) read(file *os.File) {
	partyTemplate.NumPartiesCreated.read(file)
	partyTemplate.NumPartiesDestroyed.read(file)
	partyTemplate.NumPartiesDestroyedByPlayer.read(file)
	partyTemplate.NumSlots.read(file)
	partyTemplate.Slots = make([]Int64, partyTemplate.NumSlots)
	for i := 0; i < len(partyTemplate.Slots); i++ {
		partyTemplate.Slots[i].read(file)
	}
}

func (partyStack *PartyStack) read(file *os.File) {
	partyStack.TroopId.read(file)
	partyStack.NumTroops.read(file)
	partyStack.NumWoundedTroops.read(file)
	partyStack.Flags.read(file)
}

func (party *Party) read(file *os.File, gameVersion Int32) {
	party.Id.read(file)
	party.Name.read(file)
	party.Flags.read(file)
	party.MenuId.read(file)
	party.PartyTemplateId.read(file)
	party.FactionId.read(file)
	party.Personality.read(file)
	party.DefaultBehavior.read(file)
	party.CurrentBehavior.read(file)
	party.DefaultBehaviorObjectId.read(file)
	party.CurrentBehaviorObjectId.read(file)
	party.InitialPositionX.read(file)
	party.InitialPositionY.read(file)
	party.TargetPositionX.read(file)
	party.TargetPositionY.read(file)
	party.PositionX.read(file)
	party.PositionY.read(file)
	party.PositionZ.read(file)
	party.NumStacks.read(file)
	party.Stacks = make([]PartyStack, party.NumStacks)
	for i := 0; i < len(party.Stacks); i++ {
		party.Stacks[i].read(file)
	}
	party.Bearing.read(file)
	party.Renamed.read(file)
	party.ExtraText.read(file)
	party.Morale.read(file)
	party.Hunger.read(file)
	party.Unused1.read(file)
	party.PatrolRadius.read(file)
	party.Initiative.read(file)
	party.Helpfulness.read(file)
	party.LabelVisible.read(file)
	party.BanditAttraction.read(file)
	if (gameVersion >= 900 && gameVersion < 1000) || gameVersion >= 1020 {
		party.Marshall.read(file)
	}
	party.IgnorePlayerTimer.read(file)
	party.BannerMapIconId.read(file)
	if gameVersion >= 1137 {
		party.ExtraMapIconId.read(file)
		party.ExtraMapIconUpDownDistance.read(file)
		party.ExtraMapIconUpDownFrequency.read(file)
		party.ExtraMapIconRotateFrequency.read(file)
		party.ExtraMapIconFadeFrequency.read(file)
	}
	party.AttachedToPartyId.read(file)
	if gameVersion >= 1162 {
		party.Unused2.read(file)
	}
	party.IsAttached.read(file)
	party.NumAttachedPartyIds.read(file)
	party.AttachedPartyIds = make([]Int32, party.NumAttachedPartyIds)
	for i := 0; i < len(party.AttachedPartyIds); i++ {
		party.AttachedPartyIds[i].read(file)
	}
	party.NumParticleSystemIds.read(file)
	party.ParticleSystemIds = make([]Int32, party.NumParticleSystemIds)
	for i := 0; i < len(party.ParticleSystemIds); i++ {
		party.ParticleSystemIds[i].read(file)
	}
	for i := 0; i < len(party.Notes); i++ {
		party.Notes[i].read(file)
	}
	party.NumSlots.read(file)
	party.Slots = make([]Int64, party.NumSlots)
	for i := 0; i < len(party.Slots); i++ {
		party.Slots[i].read(file)
	}
}

func (partyRecord *PartyRecord) read(file *os.File, gameVersion Int32) {
	partyRecord.Valid.read(file)
	if partyRecord.Valid == 1 {
		partyRecord.RawId.read(file)
		partyRecord.Id.read(file)
		partyRecord.Party.read(file, gameVersion)
	}
}

func (playerPartyStack *PlayerPartyStack) read(file *os.File, stackIndex int) {
	playerPartyStack.Experience.read(file)
	playerPartyStack.NumUpgradeable.read(file)
	if stackIndex < 32 {
		for i := 0; i < len(playerPartyStack.TroopDnas); i++ {
			playerPartyStack.TroopDnas[i].read(file)
		}
	}
}

func (mapEvent *MapEvent) read(file *os.File) {
	mapEvent.Unused0.read(file)
	mapEvent.Type.read(file)
	mapEvent.PositionX.read(file)
	mapEvent.PositionY.read(file)
	mapEvent.LandPositionX.read(file)
	mapEvent.LandPositionY.read(file)
	mapEvent.Unused1.read(file)
	mapEvent.Unused2.read(file)
	mapEvent.AttackerPartyId.read(file)
	mapEvent.DefenderPartyId.read(file)
	mapEvent.BattleSimulationTimer.read(file)
	mapEvent.NextBattleSimulation.read(file)
}

func (mapEventRecord *MapEventRecord) read(file *os.File) {
	mapEventRecord.Valid.read(file)
	if mapEventRecord.Valid == 1 {
		mapEventRecord.Id.read(file)
		mapEventRecord.MapEvent.read(file)
	}
}

func (item *Item) read(file *os.File) {
	item.ItemKindId.read(file)
	item.ItemFlags.read(file)
}

func (troop *Troop) read(file *os.File) {
	troop.NumSlots.read(file)
	troop.Slots = make([]Int64, troop.NumSlots)
	for i := 0; i < len(troop.Slots); i++ {
		troop.Slots[i].read(file)
	}
	for i := 0; i < len(troop.Attributes); i++ {
		troop.Attributes[i].read(file)
	}
	for i := 0; i < len(troop.Proficiencies); i++ {
		troop.Proficiencies[i].read(file)
	}
	for i := 0; i < len(troop.Skills); i++ {
		troop.Skills[i].read(file)
	}
	for i := 0; i < len(troop.Notes); i++ {
		troop.Notes[i].read(file)
	}
	troop.Flags.read(file)
	troop.SiteIdAndEntryNo.read(file)
	troop.SkillPoints.read(file)
	troop.AttributePoints.read(file)
	troop.ProficiencyPoints.read(file)
	troop.Level.read(file)
	isHero := troop.Flags&heroFlag != 0
	if isHero || loadRegularTroopInventory {
		troop.Gold.read(file)
		troop.Experience.read(file)
		troop.Health.read(file)
		troop.FactionId.read(file)
		for i := 0; i < len(troop.InventoryItems); i++ {
			troop.InventoryItems[i].read(file)
		}
		for i := 0; i < len(troop.EquippedItems); i++ {
			troop.EquippedItems[i].read(file)
		}
		for i := 0; i < len(troop.FaceKeys); i++ {
			troop.FaceKeys[i].read(file)
		}
		troop.Renamed.read(file)
		if troop.Renamed {
			troop.Name.read(file)
			troop.NamePlural.read(file)
		}
	}
	troop.ClassNo.read(file)
}

func (itemKind *ItemKind) read(file *os.File) {
	itemKind.NumSlots.read(file)
	itemKind.Slots = make([]Int64, itemKind.NumSlots)
	for i := 0; i < len(itemKind.Slots); i++ {
		itemKind.Slots[i].read(file)
	}
}

func (game *Game) read(file *os.File) {
	game.Header.read(file)
	game.GameTime.read(file)
	game.RandomSeed.read(file)
	game.SaveMode.read(file)
	if game.Header.GameVersion >= 1137 {
		game.CombatDifficulty.read(file)
		game.CombatDifficultyFriendlies.read(file)
		game.ReduceCombatAi.read(file)
		game.ReduceCampaignAi.read(file)
		game.CombatSpeed.read(file)
	}
	game.DateTimer.read(file)
	game.Hour.read(file)
	game.Day.read(file)
	game.Week.read(file)
	game.Month.read(file)
	game.Year.read(file)
	game.Unused0.read(file)
	game.GlobalCloudAmount.read(file)
	game.GlobalHazeAmount.read(file)
	game.AverageDifficulty.read(file)
	game.AverageDifficultyPeriod.read(file)
	game.Unused1.read(file)
	game.Unused2.read(file)
	game.TutorialFlags.read(file)
	game.DefaultPrisonerPrice.read(file)
	game.EncounteredParty1Id.read(file)
	game.EncounteredParty2Id.read(file)
	game.CurrentMenuId.read(file)
	game.CurrentSiteId.read(file)
	game.CurrentEntryNo.read(file)
	game.CurrentMissionTemplateId.read(file)
	game.PartyCreationMinRandomValue.read(file)
	game.PartyCreationMaxRandomValue.read(file)
	game.GameLog.read(file)
	for i := 0; i < len(game.Unused3); i++ {
		game.Unused3[i].read(file)
	}
	game.Unused4.read(file)
	game.RestPeriod.read(file)
	game.RestTimeSpeed.read(file)
	game.RestIsInteractive.read(file)
	game.RestRemainAttackable.read(file)
	for i := 0; i < len(game.ClassNames); i++ {
		game.ClassNames[i].read(file)
	}
	game.NumGlobalVariables.read(file)
	game.GlobalVariables = make([]Int64, game.NumGlobalVariables)
	for i := 0; i < len(game.GlobalVariables); i++ {
		game.GlobalVariables[i].read(file)
	}
	game.NumTriggers.read(file)
	game.Triggers = make([]Trigger, game.NumTriggers)
	for i := 0; i < len(game.Triggers); i++ {
		game.Triggers[i].read(file)
	}
	game.NumSimpleTriggers.read(file)
	game.SimpleTriggers = make([]SimpleTrigger, game.NumSimpleTriggers)
	for i := 0; i < len(game.SimpleTriggers); i++ {
		game.SimpleTriggers[i].read(file)
	}
	game.NumQuests.read(file)
	game.Quests = make([]Quest, game.NumQuests)
	for i := 0; i < len(game.Quests); i++ {
		game.Quests[i].read(file)
	}
	game.NumInfoPages.read(file)
	game.InfoPages = make([]InfoPage, game.NumInfoPages)
	for i := 0; i < len(game.InfoPages); i++ {
		game.InfoPages[i].read(file)
	}
	game.NumSites.read(file)
	game.Sites = make([]Site, game.NumSites)
	for i := 0; i < len(game.Sites); i++ {
		game.Sites[i].read(file)
	}
	game.NumFactions.read(file)
	game.Factions = make([]Faction, game.NumFactions)
	for i := 0; i < len(game.Factions); i++ {
		game.Factions[i].Relations = make([]Float, game.NumFactions)
		game.Factions[i].read(file)
	}
	game.NumMapTracks.read(file)
	game.MapTracks = make([]MapTrack, game.NumMapTracks)
	for i := 0; i < len(game.MapTracks); i++ {
		game.MapTracks[i].read(file)
	}
	game.NumPartyTemplates.read(file)
	game.PartyTemplates = make([]PartyTemplate, game.NumPartyTemplates)
	for i := 0; i < len(game.PartyTemplates); i++ {
		game.PartyTemplates[i].read(file)
	}
	game.NumPartyRecords.read(file)
	game.NumPartiesCreated.read(file)
	game.PartyRecords = make([]PartyRecord, game.NumPartyRecords)
	for i := 0; i < len(game.PartyRecords); i++ {
		game.PartyRecords[i].read(file, game.Header.GameVersion)
	}
	playerParty := game.PartyRecords[0].Party
	game.PlayerPartyStackAdditionalInfo = make([]PlayerPartyStack, len(playerParty.Stacks))
	for i := 0; i < len(playerParty.Stacks); i++ {
		if hasAdditionalInfo(playerParty, i) {
			game.PlayerPartyStackAdditionalInfo[i].read(file, i)
		}
	}
	game.NumMapEventRecords.read(file)
	game.NumMapEventsCreated.read(file)
	game.MapEventRecords = make([]MapEventRecord, game.NumMapEventRecords)
	for i := 0; i < len(game.MapEventRecords); i++ {
		game.MapEventRecords[i].read(file)
	}
	game.NumTroops.read(file)
	game.Troops = make([]Troop, game.NumTroops)
	for i := 0; i < len(game.Troops); i++ {
		game.Troops[i].read(file)
	}
	for i := 0; i < len(game.Unused5); i++ {
		game.Unused5[i].read(file)
	}
	game.NumItemKinds.read(file)
	game.ItemKinds = make([]ItemKind, game.NumItemKinds)
	for i := 0; i < len(game.ItemKinds); i++ {
		game.ItemKinds[i].read(file)
	}
	game.PlayerFaceKeys0.read(file)
	game.PlayerFaceKeys1.read(file)
	game.PlayerKillCount.read(file)
	game.PlayerWoundedCount.read(file)
	game.PlayerOwnTroopKillCount.read(file)
	game.PlayerOwnTroopWoundedCount.read(file)
}

func Load(path string) (game Game, err error) {
	file, err := os.Open(path)
	if err != nil {
		return game, err
	}
	defer file.Close()
	game.read(file)
	return game, nil
}
