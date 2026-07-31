package main

import (
	"bytes"
	"crypto/sha256"
	"io"
	"os"
	"reflect"
	"testing"
)

const (
	inPath  = "sg04.sav"
	outPath = "out.sav"
)

func TestCompareData(t *testing.T) {
	DisableNaN = true
	game, err := load(inPath)
	if err != nil {
		t.Errorf("Could not load from: %s due to error:\n%s", inPath, err)
	}
	err = save(game, outPath)
	if err != nil {
		t.Errorf("Could not save to: %s due to error:\n%s", outPath, err)
	}
	game1, err := load(outPath)
	if err != nil {
		t.Errorf("Could not load from: %s due to error:\n%s", outPath, err)
	}
	if !reflect.DeepEqual(game, game1) {
		t.Errorf("%s's data was different after saving and reloading", inPath)
	}
}

func TestCompareSaveFiles(t *testing.T) {
	DisableNaN = false
	game, err := load(inPath)
	if err != nil {
		t.Errorf("Could not load from: %s due to error:\n%s", inPath, err)
	}
	err = save(game, outPath)
	if err != nil {
		t.Errorf("Could not save to: %s due to error:\n%s", outPath, err)
	}
	h1, err := hashFile(inPath)
	if err != nil {
		t.Errorf("Could not hash %s due to error:\n%s", outPath, err)
	}
	h2, err := hashFile(outPath)
	if err != nil {
		t.Errorf("Could not hash %s due to error:\n%s", outPath, err)
	}
	if !bytes.Equal(h1, h2) {
		t.Errorf("%s's hash was different after loading and saving", inPath)
	}
}

func hashFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	hasher := sha256.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}
