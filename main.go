package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var JSON_DEBUG = false

func load(path string) (game Game, err error) {
	file, err := os.Open(path)
	if err != nil {
		return Game{}, err
	}
	defer file.Close()

	game.Read(file)

	return
}

func save(path string, game Game) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf, err := game.Write()
	if err != nil {
		return err
	}

	return binary.Write(file, binary.LittleEndian, buf)
}

func main() {
	inPath := "C:/Users/Daniel/Documents/Mount&Blade Warband Savegames/Vexed Native 1.154/sg04.sav"

	JSON_DEBUG = true

	game, err := load(inPath)
	if err != nil {
		log.Fatal(err)
	}

	ExportToJSON("out.json", game)

	err = save("out.sav", game)
	if err != nil {
		log.Fatal(err)
	}

	game1, err := load("out.sav")
	if err != nil {
		log.Fatal(err)
	}

	ExportToJSON("out1.json", game1)
}

func ExportToJSON(path string, game Game) {
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

func PrintJSON(gameObject any) {
	bytes, err := json.MarshalIndent(gameObject, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
