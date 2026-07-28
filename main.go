package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

var JsonDebug = false

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
	JsonDebug = true
	game, err := load(inPath)
	if err != nil {
		panic(err)
	}
	ExportToJSON("out.json", game)
	err = save("out.sav", game)
	if err != nil {
		panic(err)
	}
	game1, err := load("out.sav")
	if err != nil {
		panic(err)
	}
	ExportToJSON("out1.json", game1)
	/*game, err := load("sg03.sav")
	if err != nil {
		panic(err)
	}
	err = save("out3.sav", game)
	if err != nil {
		panic(err)
	}
	game, err = load("sg04.sav")
	if err != nil {
		panic(err)
	}
	err = save("out4.sav", game)
	if err != nil {
		panic(err)
	}
	game, err = load("sg05.sav")
	if err != nil {
		panic(err)
	}
	err = save("out5.sav", game)
	if err != nil {
		panic(err)
	}*/
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
