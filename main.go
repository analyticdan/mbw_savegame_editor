package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"os"
)

const JSON_DEBUG = false

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

	outPath := "out.sav"

	game, err := load(inPath)
	if err != nil {
		log.Fatal(err)
	}

	if JSON_DEBUG {
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
	}

	err = save(outPath, game)
	if err != nil {
		log.Fatal(err)
	}

	if JSON_DEBUG {
		game, err := load(inPath)
		if err != nil {
			log.Fatal(err)
		}
		out, err := os.Create("out1.json")
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

	//bytes, _ := json.MarshalIndent(game.PartyRecords[0].Party.Id, "", "  ")
	//fmt.Println(string(bytes))

	//fmt.Println(game.PlayerFaceKeys0, game.PlayerFaceKeys1)
}
