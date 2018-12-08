package main

import (
	"os"
	"testing"
)

// test size of a new deck
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "As de Coeur" {
		t.Errorf("Expected As de Coeur, but got %v", d[0])
	}

	if d[len(d)-1] != "Roi de Trefle" {
		t.Errorf("Expected Roi de trefle, but got %v", d[len(d)-1])
	}
}

// test read and write of a file
func TestSaveAndReadFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	load := newDeckFromFile(filename)

	if len(load) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	os.Remove(filename)
}
