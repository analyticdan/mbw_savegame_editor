package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func load(path string) (game Game, err error) {
	file, err := os.Open(path)
	if err != nil {
		return Game{}, err
	}
	defer file.Close()

	game.Read(file)

	return
}

func main() {
	path := "C:/Users/Daniel/Documents/Mount&Blade Warband Savegames/Vexed Native 1.154/sg03.sav"

	game, err := load(path)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("out.json")
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

	/*bytes, _ := json.MarshalIndent(game.ItemKinds, "", "  ")
	fmt.Println(string(bytes))*/

	fmt.Println(game.PlayerFaceKeys0, game.PlayerFaceKeys1)
}
