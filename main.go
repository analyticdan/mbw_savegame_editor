package main

import (
	"cmp"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

var DisableNaN = false

func load(path string) (game Game, err error) {
	file, err := os.Open(path)
	if err != nil {
		return Game{}, err
	}
	defer file.Close()
	game.Read(file)
	return game, nil
}

func save(path string, game Game) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := game.Write(nil)
	return binary.Write(file, binary.LittleEndian, buf)
}

func main() {
	inPath := "C:/Users/Daniel/Documents/Mount&Blade Warband Savegames/Vexed Native 1.154/sg06.sav"
	game, err := load(inPath)
	if err != nil {
		panic(err)
	}

	printNegativeReputationFiefs(game)
	printCompanionLocations(game)
	printVillagesInfestedByBandits(game, false)
	printTownsAndCastlesByGarrisonSize(game)

}

func printNegativeReputationFiefs(game Game) {
	fmt.Println("Negative reputation fiefs: ")
	for _, negativeReputationFief := range getNegativeReputationFiefs(game) {
		fiefName := negativeReputationFief.Name
		fiefReputation := getFiefReputation(negativeReputationFief)
		fmt.Printf("%s, Reputation: %d\n", fiefName, fiefReputation)
	}
	fmt.Println("---")
}

func printCompanionLocations(game Game) {
	fmt.Println("Companion locations (does not include imprisoned companions or companions on a mission): ")
	for companionId, companionName := range Companions {
		companion := game.Troops[companionId]
		locationId := getCompanionLocationId(companion)
		if locationId != -1 {
			locationName := game.PartyRecords[locationId].Party.Name
			fmt.Printf("%s: %s\n", companionName, locationName)
		}
	}
	fmt.Println("---")
}

func printVillagesInfestedByBandits(game Game, includeLootedVillages bool) {
	if includeLootedVillages {
		fmt.Println("Villages infested by bandits (includes looted villages):")
	} else {
		fmt.Println("Villages infested by bandits (does not include looted villages):")
	}
	for _, village := range getVillagesInfestedByBandits(game) {
		if includeLootedVillages || isVillageStateNormal(village) {
			townOrCastle := game.PartyRecords[getVillageTownOrCastleId(village)].Party
			market := game.PartyRecords[getVillageMarketId(village)].Party
			villageName := village.Name
			townOrCastleName := townOrCastle.Name
			marketName := market.Name
			fmt.Printf("%s (Associated Castle or Town: %s, Associated Market: %s)\n", villageName, townOrCastleName, marketName)
		}
	}
	fmt.Println("---")
}

func printTownsAndCastlesByGarrisonSize(game Game) {
	fmt.Println("Towns and Castles by Garrison Size:")
	townsAndCastles := getTownsAndCastles(game)
	slices.SortFunc(townsAndCastles, func(a, b Party) int {
		return cmp.Compare(getGarrisonSize(a), getGarrisonSize(b))
	})
	for _, townOrCastle := range townsAndCastles {
		name := townOrCastle.Name
		size := getGarrisonSize(townOrCastle)
		fmt.Printf("%s: %d troops\n", name, size)
	}
	fmt.Println("---")
}

func ExportToJson(path string, game Game) {
	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(game)
	if err != nil {
		panic(err)
	}
}

func PrintJson(gameObject any) {
	bytes, err := json.MarshalIndent(gameObject, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
