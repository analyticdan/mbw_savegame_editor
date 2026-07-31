package main

import (
	"cmp"
	"fmt"
	. "mbw-savegame-editor/savegame/savegame"
	"slices"
)

func main() {
	inPath := "C:/Users/Daniel/Documents/Mount&Blade Warband Savegames/Vexed Native 1.154/sg06.sav"
	game, err := Load(inPath)
	if err != nil {
		panic(err)
	}

	//printCompanionsByPolearmProficiency(game)
	//printCompanionLocations(game)
	//printNegativeReputationFiefs(game)
	printVillagesInfestedByBandits(game)
	printFortificationsByGarrisonSize(game, KhergitKhanate)
	printFortificationsByGarrisonSize(game, SarranidSultanate)
}

func printNegativeReputationFiefs(game Game) {
	fmt.Println("Negative reputation fiefs: ")
	for _, fief := range getAllFiefs(game) {
		if reputation := getFiefReputation(fief); reputation < 0 {
			if isVillage(fief) {
				fortification := getVillageFortification(game, fief)
				market := getVillageMarket(game, fief)
				fmt.Printf("%s (fortification: %s, market: %s), reputation: %d\n", fief.Name, &fortification.Name, market.Name, reputation)
			} else {
				fmt.Printf("%s; reputation: %d\n", fief.Name, reputation)
			}
		}
	}
	fmt.Println("---")
}

func printCompanionLocations(game Game) {
	fmt.Println("Companion locations (does not include imprisoned companions or companions on a mission):")
	for _, companionId := range CompanionIds {
		companion := getTroop(game, companionId)
		if locationId := getTroopLocationId(companion); locationId != -1 {
			location := getFief(game, locationId)
			fmt.Printf("%s: %s\n", CompanionsNameMap[companionId], location.Name)
		}
	}
	fmt.Println("---")
}

func printCompanionsByPolearmProficiency(game Game) {
	fmt.Println("Companions by proficiencies:")
	companionIds := CompanionIds
	getPolearmProficiencies := func(troop Troop) Float {
		return troop.Proficiencies[2]
	}
	slices.SortFunc(companionIds, func(a, b int) int {
		aValue := getPolearmProficiencies(getTroop(game, a))
		bValue := getPolearmProficiencies(getTroop(game, b))
		return cmp.Compare(aValue, bValue)
	})
	for _, companionId := range companionIds {
		companion := getTroop(game, companionId)
		fmt.Printf("%s: %.0f\n", CompanionsNameMap[companionId], getPolearmProficiencies(companion))
	}
	fmt.Println("---")
}

func printVillagesInfestedByBandits(game Game) {
	fmt.Println("Villages infested by bandits:")
	var status string
	for _, village := range getVillages(game) {
		if isVillageInfestedByBandits(village) {
			fortification := getVillageFortification(game, village)
			market := getVillageMarket(game, village)
			if state := getVillageState(village); state == VillageBeingRaided {
				status = " (Being Raided) "
			} else if state == VillageLooted {
				status = " (Looted) "
			}
			fmt.Printf("%s%s (fortification: %s, market: %s)\n", village.Name, status, fortification.Name, market.Name)
		}
	}
	fmt.Println("---")
}

func printFortificationsByGarrisonSize(game Game, factionId int) {
	fmt.Printf("%s's Fortifications by Garrison Size:\n", getFaction(game, factionId).Name)
	fortifications := getFortifications(game)
	slices.SortFunc(fortifications, func(a, b Party) int {
		return cmp.Compare(getGarrisonSize(a), getGarrisonSize(b))
	})
	for _, fortification := range fortifications {

		if factionId == int(fortification.FactionId) {
			var ladderString string
			if isFortificationSiegedWithLadders(fortification) {
				ladderString = " (ladder)"
			}
			fmt.Printf("%s%s: %d troops\n", fortification.Name, ladderString, getGarrisonSize(fortification))
		}
	}
	fmt.Println("---")
}

func unequipCompanionItems(game Game, equipmentSlot int, inventOffset int) {
	// Use order of proficiencies above to ensure the best characters get the first items.
	heroIds := []int{197, 199, 203, 202, 207, 201, 198, 200, 206, 208, 209, 204, 194, 195, 205, 196}

	for i, heroId := range heroIds {
		game.Troops[0].InventoryItems[inventOffset+i] = game.Troops[heroId].EquippedItems[equipmentSlot]
		game.Troops[heroId].EquippedItems[equipmentSlot].ItemKindId = -1
	}
}

func equipCompanionItems(game Game, equipmentSlot int, inventOffset int) {
	// Use order of proficiencies above to ensure the best characters get the first items.
	heroIds := []int{197, 199, 203, 202, 207, 201, 198, 200, 206, 208, 209, 204, 194, 195, 205, 196}

	for i, heroId := range heroIds {
		game.Troops[heroId].EquippedItems[equipmentSlot] = game.Troops[0].InventoryItems[inventOffset+i]
		game.Troops[0].InventoryItems[inventOffset+i].ItemKindId = -1
	}
}
