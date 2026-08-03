package savegame

import (
	"encoding/binary"
	"os"
)

func (b *Bool) append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, b)
	if err != nil {
		panic(err)
	}
	return buf
}

func (i *Int32) append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
	return buf
}

func (i *Int64) append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
	return buf
}

func (i *UInt32) append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
	return buf
}

func (i *UInt64) append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
	return buf
}

func (f *Float) append(buf []byte) []byte {
	buf, err := binary.Append(buf, binary.LittleEndian, f)
	if err != nil {
		panic(err)
	}
	return buf
}

func (s *String) append(buf []byte) []byte {
	buf = s.NumChars.append(buf)
	buf, err := binary.Append(buf, binary.LittleEndian, s.Chars)
	if err != nil {
		panic(err)
	}
	return buf
}

func (header *Header) append(buf []byte) []byte {
	buf = header.MagicNumber.append(buf)
	buf = header.GameVersion.append(buf)
	buf = header.ModuleVersion.append(buf)
	buf = header.SavegameName.append(buf)
	buf = header.PlayerName.append(buf)
	buf = header.PlayerLevel.append(buf)
	buf = header.Date.append(buf)
	return buf
}

func (trigger *Trigger) append(buf []byte) []byte {
	buf = trigger.Status.append(buf)
	buf = trigger.CheckTimer.append(buf)
	buf = trigger.DelayTimer.append(buf)
	buf = trigger.RearmTimer.append(buf)
	return buf
}

func (simpleTrigger *SimpleTrigger) append(buf []byte) []byte {
	buf = simpleTrigger.CheckTimer.append(buf)
	return buf
}

func (note *Note) append(buf []byte) []byte {
	buf = note.Text.append(buf)
	buf = note.Value.append(buf)
	buf = note.TableauMaterialId.append(buf)
	buf = note.Available.append(buf)
	return buf
}

func (quest *Quest) append(buf []byte) []byte {
	buf = quest.Progression.append(buf)
	buf = quest.GiverTroopId.append(buf)
	buf = quest.Number.append(buf)
	buf = quest.StartDate.append(buf)
	buf = quest.Title.append(buf)
	buf = quest.Text.append(buf)
	buf = quest.Giver.append(buf)
	for i := 0; i < len(quest.Notes); i++ {
		buf = quest.Notes[i].append(buf)
	}
	buf = quest.NumSlots.append(buf)
	for i := 0; i < len(quest.Slots); i++ {
		buf = quest.Slots[i].append(buf)
	}
	return buf
}

func (infoPage *InfoPage) append(buf []byte) []byte {
	for i := 0; i < len(infoPage.Notes); i++ {
		buf = infoPage.Notes[i].append(buf)
	}
	return buf
}

func (site *Site) append(buf []byte) []byte {
	buf = site.NumSlots.append(buf)
	for i := 0; i < len(site.Slots); i++ {
		buf = site.Slots[i].append(buf)
	}
	return buf
}

func (faction *Faction) append(buf []byte) []byte {
	buf = faction.NumSlots.append(buf)
	for i := 0; i < len(faction.Slots); i++ {
		buf = faction.Slots[i].append(buf)
	}
	for i := 0; i < len(faction.Relations); i++ {
		buf = faction.Relations[i].append(buf)
	}
	buf = faction.Name.append(buf)
	buf = faction.Renamed.append(buf)
	buf = faction.Color.append(buf)
	buf = faction.Unused.append(buf)
	for i := 0; i < len(faction.Notes); i++ {
		buf = faction.Notes[i].append(buf)
	}
	return buf
}

func (mapTrack *MapTrack) append(buf []byte) []byte {
	buf = mapTrack.PositionX.append(buf)
	buf = mapTrack.PositionY.append(buf)
	buf = mapTrack.PositionZ.append(buf)
	buf = mapTrack.Rotation.append(buf)
	buf = mapTrack.Age.append(buf)
	buf = mapTrack.Flags.append(buf)
	return buf
}

func (partyTemplate *PartyTemplate) append(buf []byte) []byte {
	buf = partyTemplate.NumPartiesCreated.append(buf)
	buf = partyTemplate.NumPartiesDestroyed.append(buf)
	buf = partyTemplate.NumPartiesDestroyedByPlayer.append(buf)
	buf = partyTemplate.NumSlots.append(buf)
	for i := 0; i < len(partyTemplate.Slots); i++ {
		buf = partyTemplate.Slots[i].append(buf)
	}
	return buf
}

func (partyStack *PartyStack) append(buf []byte) []byte {
	buf = partyStack.TroopId.append(buf)
	buf = partyStack.NumTroops.append(buf)
	buf = partyStack.NumWoundedTroops.append(buf)
	buf = partyStack.Flags.append(buf)
	return buf
}

func (party *Party) append(buf []byte, gameVersion Int32) []byte {
	buf = party.Id.append(buf)
	buf = party.Name.append(buf)
	buf = party.Flags.append(buf)
	buf = party.MenuId.append(buf)
	buf = party.PartyTemplateId.append(buf)
	buf = party.FactionId.append(buf)
	buf = party.Personality.append(buf)
	buf = party.DefaultBehavior.append(buf)
	buf = party.CurrentBehavior.append(buf)
	buf = party.DefaultBehaviorObjectId.append(buf)
	buf = party.CurrentBehaviorObjectId.append(buf)
	buf = party.InitialPositionX.append(buf)
	buf = party.InitialPositionY.append(buf)
	buf = party.TargetPositionX.append(buf)
	buf = party.TargetPositionY.append(buf)
	buf = party.PositionX.append(buf)
	buf = party.PositionY.append(buf)
	buf = party.PositionZ.append(buf)
	buf = party.NumStacks.append(buf)
	for i := 0; i < len(party.Stacks); i++ {
		buf = party.Stacks[i].append(buf)
	}
	buf = party.Bearing.append(buf)
	buf = party.Renamed.append(buf)
	buf = party.ExtraText.append(buf)
	buf = party.Morale.append(buf)
	buf = party.Hunger.append(buf)
	buf = party.Unused1.append(buf)
	buf = party.PatrolRadius.append(buf)
	buf = party.Initiative.append(buf)
	buf = party.Helpfulness.append(buf)
	buf = party.LabelVisible.append(buf)
	buf = party.BanditAttraction.append(buf)
	if (gameVersion >= 900 && gameVersion < 1000) || gameVersion >= 1020 {
		buf = party.Marshall.append(buf)
	}
	buf = party.IgnorePlayerTimer.append(buf)
	buf = party.BannerMapIconId.append(buf)
	if gameVersion >= 1137 {
		buf = party.ExtraMapIconId.append(buf)
		buf = party.ExtraMapIconUpDownDistance.append(buf)
		buf = party.ExtraMapIconUpDownFrequency.append(buf)
		buf = party.ExtraMapIconRotateFrequency.append(buf)
		buf = party.ExtraMapIconFadeFrequency.append(buf)
	}
	buf = party.AttachedToPartyId.append(buf)
	if gameVersion >= 1162 {
		buf = party.Unused2.append(buf)
	}
	buf = party.IsAttached.append(buf)
	buf = party.NumAttachedPartyIds.append(buf)
	for i := 0; i < len(party.AttachedPartyIds); i++ {
		buf = party.AttachedPartyIds[i].append(buf)
	}
	buf = party.NumParticleSystemIds.append(buf)
	for i := 0; i < len(party.ParticleSystemIds); i++ {
		buf = party.ParticleSystemIds[i].append(buf)
	}
	for i := 0; i < len(party.Notes); i++ {
		buf = party.Notes[i].append(buf)
	}
	buf = party.NumSlots.append(buf)
	for i := 0; i < len(party.Slots); i++ {
		buf = party.Slots[i].append(buf)
	}
	return buf
}

func (partyRecord *PartyRecord) append(buf []byte, gameVersion Int32) []byte {
	buf = partyRecord.Valid.append(buf)
	if partyRecord.Valid == 1 {
		buf = partyRecord.RawId.append(buf)
		buf = partyRecord.Id.append(buf)
		buf = partyRecord.Party.append(buf, gameVersion)
	}
	return buf
}

func (playerPartyStack *PlayerPartyStack) append(buf []byte, stackIndex int) []byte {
	buf = playerPartyStack.Experience.append(buf)
	buf = playerPartyStack.NumUpgradeable.append(buf)
	if stackIndex < 32 {
		for i := 0; i < len(playerPartyStack.TroopDnas); i++ {
			buf = playerPartyStack.TroopDnas[i].append(buf)
		}
	}
	return buf
}

func (mapEvent *MapEvent) append(buf []byte) []byte {
	buf = mapEvent.Unused0.append(buf)
	buf = mapEvent.Type.append(buf)
	buf = mapEvent.PositionX.append(buf)
	buf = mapEvent.PositionY.append(buf)
	buf = mapEvent.LandPositionX.append(buf)
	buf = mapEvent.LandPositionY.append(buf)
	buf = mapEvent.Unused1.append(buf)
	buf = mapEvent.Unused2.append(buf)
	buf = mapEvent.AttackerPartyId.append(buf)
	buf = mapEvent.DefenderPartyId.append(buf)
	buf = mapEvent.BattleSimulationTimer.append(buf)
	buf = mapEvent.NextBattleSimulation.append(buf)
	return buf
}

func (mapEventRecord *MapEventRecord) append(buf []byte) []byte {
	buf = mapEventRecord.Valid.append(buf)
	if mapEventRecord.Valid == 1 {
		buf = mapEventRecord.Id.append(buf)
		buf = mapEventRecord.MapEvent.append(buf)
	}
	return buf
}

func (item *Item) append(buf []byte) []byte {
	buf = item.ItemKindId.append(buf)
	buf = item.ItemFlags.append(buf)
	return buf
}

func (troop *Troop) append(buf []byte) []byte {
	buf = troop.NumSlots.append(buf)
	for i := 0; i < len(troop.Slots); i++ {
		buf = troop.Slots[i].append(buf)
	}
	for i := 0; i < len(troop.Attributes); i++ {
		buf = troop.Attributes[i].append(buf)
	}
	for i := 0; i < len(troop.Proficiencies); i++ {
		buf = troop.Proficiencies[i].append(buf)
	}
	for i := 0; i < len(troop.Skills); i++ {
		buf = troop.Skills[i].append(buf)
	}
	for i := 0; i < len(troop.Notes); i++ {
		buf = troop.Notes[i].append(buf)
	}
	buf = troop.Flags.append(buf)
	buf = troop.SiteIdAndEntryNo.append(buf)
	buf = troop.SkillPoints.append(buf)
	buf = troop.AttributePoints.append(buf)
	buf = troop.ProficiencyPoints.append(buf)
	buf = troop.Level.append(buf)
	isHero := troop.Flags&heroFlag != 0
	if isHero || loadRegularTroopInventory {
		buf = troop.Gold.append(buf)
		buf = troop.Experience.append(buf)
		buf = troop.Health.append(buf)
		buf = troop.FactionId.append(buf)
		for i := 0; i < len(troop.InventoryItems); i++ {
			buf = troop.InventoryItems[i].append(buf)
		}
		for i := 0; i < len(troop.EquippedItems); i++ {
			buf = troop.EquippedItems[i].append(buf)
		}
		for i := 0; i < len(troop.FaceKeys); i++ {
			buf = troop.FaceKeys[i].append(buf)
		}
		buf = troop.Renamed.append(buf)
		if troop.Renamed {
			buf = troop.Name.append(buf)
			buf = troop.NamePlural.append(buf)
		}
	}
	buf = troop.ClassNo.append(buf)
	return buf
}

func (itemKind *ItemKind) append(buf []byte) []byte {
	buf = itemKind.NumSlots.append(buf)
	for i := 0; i < len(itemKind.Slots); i++ {
		buf = itemKind.Slots[i].append(buf)
	}
	return buf
}

func (game *Game) write() []byte {
	var buf []byte
	buf = game.Header.append(buf)
	buf = game.GameTime.append(buf)
	buf = game.RandomSeed.append(buf)
	buf = game.SaveMode.append(buf)
	if game.Header.GameVersion >= 1137 {
		buf = game.CombatDifficulty.append(buf)
		buf = game.CombatDifficultyFriendlies.append(buf)
		buf = game.ReduceCombatAi.append(buf)
		buf = game.ReduceCampaignAi.append(buf)
		buf = game.CombatSpeed.append(buf)
	}
	buf = game.DateTimer.append(buf)
	buf = game.Hour.append(buf)
	buf = game.Day.append(buf)
	buf = game.Week.append(buf)
	buf = game.Month.append(buf)
	buf = game.Year.append(buf)
	buf = game.Unused0.append(buf)
	buf = game.GlobalCloudAmount.append(buf)
	buf = game.GlobalHazeAmount.append(buf)
	buf = game.AverageDifficulty.append(buf)
	buf = game.AverageDifficultyPeriod.append(buf)
	buf = game.Unused1.append(buf)
	buf = game.Unused2.append(buf)
	buf = game.TutorialFlags.append(buf)
	buf = game.DefaultPrisonerPrice.append(buf)
	buf = game.EncounteredParty1Id.append(buf)
	buf = game.EncounteredParty2Id.append(buf)
	buf = game.CurrentMenuId.append(buf)
	buf = game.CurrentSiteId.append(buf)
	buf = game.CurrentEntryNo.append(buf)
	buf = game.CurrentMissionTemplateId.append(buf)
	buf = game.PartyCreationMinRandomValue.append(buf)
	buf = game.PartyCreationMaxRandomValue.append(buf)
	buf = game.GameLog.append(buf)
	for i := 0; i < len(game.Unused3); i++ {
		buf = game.Unused3[i].append(buf)
	}
	buf = game.Unused4.append(buf)
	buf = game.RestPeriod.append(buf)
	buf = game.RestTimeSpeed.append(buf)
	buf = game.RestIsInteractive.append(buf)
	buf = game.RestRemainAttackable.append(buf)
	for i := 0; i < len(game.ClassNames); i++ {
		buf = game.ClassNames[i].append(buf)
	}
	buf = game.NumGlobalVariables.append(buf)
	for i := 0; i < len(game.GlobalVariables); i++ {
		buf = game.GlobalVariables[i].append(buf)
	}
	buf = game.NumTriggers.append(buf)
	for i := 0; i < len(game.Triggers); i++ {
		buf = game.Triggers[i].append(buf)
	}
	buf = game.NumSimpleTriggers.append(buf)
	for i := 0; i < len(game.SimpleTriggers); i++ {
		buf = game.SimpleTriggers[i].append(buf)
	}
	buf = game.NumQuests.append(buf)
	for i := 0; i < len(game.Quests); i++ {
		buf = game.Quests[i].append(buf)
	}
	buf = game.NumInfoPages.append(buf)
	for i := 0; i < len(game.InfoPages); i++ {
		buf = game.InfoPages[i].append(buf)
	}
	buf = game.NumSites.append(buf)
	for i := 0; i < len(game.Sites); i++ {
		buf = game.Sites[i].append(buf)
	}
	buf = game.NumFactions.append(buf)
	for i := 0; i < len(game.Factions); i++ {
		buf = game.Factions[i].append(buf)
	}
	buf = game.NumMapTracks.append(buf)
	for i := 0; i < len(game.MapTracks); i++ {
		buf = game.MapTracks[i].append(buf)
	}
	buf = game.NumPartyTemplates.append(buf)
	for i := 0; i < len(game.PartyTemplates); i++ {
		buf = game.PartyTemplates[i].append(buf)
	}
	buf = game.NumPartyRecords.append(buf)
	buf = game.NumPartiesCreated.append(buf)
	for i := 0; i < len(game.PartyRecords); i++ {
		buf = game.PartyRecords[i].append(buf, game.Header.GameVersion)
	}
	playerParty := game.PartyRecords[0].Party
	for i := 0; i < len(game.PlayerPartyStackAdditionalInfo); i++ {
		if hasAdditionalInfo(playerParty, i) {
			buf = game.PlayerPartyStackAdditionalInfo[i].append(buf, i)
		}
	}
	buf = game.NumMapEventRecords.append(buf)
	buf = game.NumMapEventsCreated.append(buf)
	for i := 0; i < len(game.MapEventRecords); i++ {
		buf = game.MapEventRecords[i].append(buf)
	}
	buf = game.NumTroops.append(buf)
	for i := 0; i < len(game.Troops); i++ {
		buf = game.Troops[i].append(buf)
	}
	for i := 0; i < len(game.Unused5); i++ {
		buf = game.Unused5[i].append(buf)
	}
	buf = game.NumItemKinds.append(buf)
	for i := 0; i < len(game.ItemKinds); i++ {
		buf = game.ItemKinds[i].append(buf)
	}
	buf = game.PlayerFaceKeys0.append(buf)
	buf = game.PlayerFaceKeys1.append(buf)
	buf = game.PlayerKillCount.append(buf)
	buf = game.PlayerWoundedCount.append(buf)
	buf = game.PlayerOwnTroopKillCount.append(buf)
	buf = game.PlayerOwnTroopWoundedCount.append(buf)
	return buf
}

func Save(game Game, path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := game.write()
	return binary.Write(file, binary.LittleEndian, buf)
}
