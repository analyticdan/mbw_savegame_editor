package main

import (
	. "mbw-savegame-editor/savegame/savegame"
	"slices"
)

func getFief(game Game, fiefId int) Party {
	return game.PartyRecords[fiefId].Party
}

func getFiefs(game Game, fiefIds []int) []Party {
	fiefs := make([]Party, len(fiefIds))
	for i, fiefId := range fiefIds {
		fiefs[i] = getFief(game, fiefId)
	}
	return fiefs
}

func getTowns(game Game) []Party {
	return getFiefs(game, TownIds)
}

func getCastles(game Game) []Party {
	return getFiefs(game, CastleIds)
}

func getVillages(game Game) []Party {
	return getFiefs(game, VillageIds)
}

func getFortifications(game Game) []Party {
	return slices.Concat(getTowns(game), getCastles(game))
}

func getAllFiefs(game Game) []Party {
	return slices.Concat(getTowns(game), getCastles(game), getVillages(game))
}

func getFiefReputation(fief Party) int {
	return int(fief.Slots[26])
}

func getFiefOriginalFactionId(fief Party) int {
	return int(fief.Slots[61])
}

func isVillage(party Party) bool {
	return party.Slots[0] == 4
}

func getVillageState(village Party) int {
	return int(village.Slots[35])
}

func isVillageInfestedByBandits(village Party) bool {
	return village.Slots[39] != 0
}

func getVillageFortification(game Game, village Party) Party {
	return getFief(game, int(village.Slots[120]))
}

func getVillageMarket(game Game, village Party) Party {
	return getFief(game, int(village.Slots[121]))
}

func hasTownEnterprise(town Party) bool {
	// cf. center_player_enterprise module_constants.py
	return town.Slots[137] != 0

}

func isFortificationSiegedWithLadders(fortification Party) bool {
	// cf. slot_center_siege_with_belfry in module_constants.py
	return fortification.Slots[27] == 0
}

func getGarrisonSize(party Party) int {
	size := 0
	for _, stack := range party.Stacks {
		if stack.Flags == 0 {
			size += int(stack.NumTroops)
		}
	}
	return size
}

func getTroop(game Game, troopId int) Troop {
	return game.Troops[troopId]
}

func getTroopRenown(troop Troop) int {
	return int(troop.Slots[7])
}

func getTroopLocationId(troop Troop) int {
	return int(troop.Slots[12])
}

func getFaction(game Game, factionId int) Faction {
	return game.Factions[factionId]
}

//Lord reputation = Troop.Slots[52]
