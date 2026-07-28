package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
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
	inPath := "sg03.sav"
	game, err := load(inPath)
	if err != nil {
		panic(err)
	}

	for _, faction := range game.Factions {
		PrintJson(faction.Name)
	}
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
