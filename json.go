package main

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/analyticdan/mbw-savegame-editor/data"
)

func ExportToJson(game Game, path string) {
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
