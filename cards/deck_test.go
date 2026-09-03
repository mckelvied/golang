package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)

	d := newDeck()
	err := d.saveToFile(filename)

	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}

	loadedDeck := newDeckFromFile(filename)

	// check deck lengths match
	if len(loadedDeck) != len(d) {
		t.Errorf("Expected deck length of %v, but got %v", len(d), len(loadedDeck))
	}

	// check first and last card match
	if loadedDeck[0] != d[0] {
		t.Errorf("Expected first card to be %v, but got %v", d[0], loadedDeck[0])
	}
	if loadedDeck[len(loadedDeck)-1] != d[len(d)-1] {
		t.Errorf("Expected last card to be %v, but got %v", d[len(d)-1], loadedDeck[len(loadedDeck)-1])
	}

	os.Remove(filename)
}
