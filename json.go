package main

import (
	"encoding/json"
	"fmt"
	"mbw-savegame-editor/savegame/savegame"
	"os"
)

func ExportToJson(game savegame.Game, path string) {
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
