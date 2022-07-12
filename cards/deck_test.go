package main

import "testing"

func TestNewDeck(t *testing.T){
	d:= newDeck()
	if len(d) != 20 {
		t.Errorf("expected deck length of 20, but got %d", len(d))
	}
}