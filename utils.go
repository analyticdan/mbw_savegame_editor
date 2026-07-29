package main

import "slices"

func getFiefs(game Game, fiefIds []int) []Party {
	fiefs := make([]Party, len(fiefIds))
	for i, fiefId := range fiefIds {
		fiefs[i] = game.PartyRecords[fiefId].Party
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

func getTownsAndCastles(game Game) []Party {
	return slices.Concat(getTowns(game), getCastles(game))
}

func getAllFiefs(game Game) []Party {
	return slices.Concat(getTowns(game), getCastles(game), getVillages(game))
}

func getFiefReputation(fief Party) Int64 {
	return fief.Slots[26]
}

func getNegativeReputationFiefs(game Game) []Party {
	var negativeReputationFiefs []Party
	for _, fief := range getAllFiefs(game) {
		if getFiefReputation(fief) < 0 {
			negativeReputationFiefs = append(negativeReputationFiefs, fief)
		}
	}
	return negativeReputationFiefs
}

func getVillageState(village Party) Int64 {
	return village.Slots[35]
}

func isVillageStateNormal(village Party) bool {
	return getVillageState(village) == 0
}

func isVillageInfestedByBandits(village Party) bool {
	return village.Slots[39] != 0
}

func getVillageTownOrCastleId(village Party) Int64 {
	return village.Slots[120]
}

func getVillageMarketId(village Party) Int64 {
	return village.Slots[121]
}

func getVillagesInfestedByBandits(game Game) []Party {
	var villagesInfestedByBandits []Party
	for _, village := range getVillages(game) {
		if isVillageInfestedByBandits(village) {
			villagesInfestedByBandits = append(villagesInfestedByBandits, village)
		}
	}
	return villagesInfestedByBandits
}

func getGarrisonSize(party Party) int {
	size := 0
	for _, stack := range party.Stacks {
		size += int(stack.NumTroops)
	}
	return size
}

func getCompanionLocationId(companion Troop) Int64 {
	return companion.Slots[12]
}
