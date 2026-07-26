package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	path := "/home/daniel/.mbwarband/Savegames/Native/sg00.sav"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	game := Game{}
	game.Read(file)

	out, err := os.Create("out.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "    ") 

	err = encoder.Encode(game)
	if err != nil {
		panic(err)
	}

	/*bytes, _ := json.MarshalIndent(game.PartyRecords[0], "", "  ")
	fmt.Println(string(bytes))*/

	fmt.Println(game.NumMapEventRecords)
}
