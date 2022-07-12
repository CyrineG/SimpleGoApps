package main

import (
	"errors"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){
	d:= newDeck()
	if len(d) != 16 {
		t.Errorf("expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "ace of spades" {
		t.Errorf("expected d[0] equals 'ace of spades', but got %s", d[0])
	}

	if d[len(d)-1] != "four of clubs" {
		t.Errorf("expected d[0] equals 'four of clubs', but got %s", d[len(d)-1])
	}
}


func TestSaveToFileAndDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	if _, err := os.Stat("_decktesting"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected 'decktesting' to be created, but file doesn't exist")
	}

	dfile := newDeckFromFile("_decktesting")
	if len(dfile) != 16 {
		t.Errorf("expected deck length of 16, but got %d", len(d))
	}

	os.Remove("_decktesting")
}