package savegame

import (
	"encoding/binary"
	"os"
)

func Load(path string) (game Game, err error) {
	file, err := os.Open(path)
	if err != nil {
		return Game{}, err
	}
	defer file.Close()
	game.Read(file)
	return game, nil
}

func Save(game Game, path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := game.Write(nil)
	return binary.Write(file, binary.LittleEndian, buf)
}
