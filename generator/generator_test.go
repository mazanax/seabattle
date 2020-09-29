package generator

import "testing"

func TestShipCouldBePlacedHere(t *testing.T) {
	var field string

	/*
		_ _ X _ _ X X X X _
		_ _ _ _ _ _ _ _ _ _
		_ _ _ _ _ _ X X _ _
		_ X X _ _ _ _ _ _ X
		_ _ _ _ _ _ _ _ _ _
		_ _ _ _ _ _ _ _ _ _
		_ _ _ X X X _ _ _ _
		X X _ _ _ _ _ _ _ _
		_ _ _ _ _ _ _ _ _ _
		X _ _ _ _ _ X X X _
	*/

	field="  O  OOOO                 OO   OO      O                       OOO    OO                  O     OOO "
	if shipCouldBePlacedHere([]rune(field), 2, 1, false) {
		t.Errorf("Should be false")
	}

	if shipCouldBePlacedHere([]rune(field), 16, 1, false) {
		t.Errorf("Should be false")
	}

	if shipCouldBePlacedHere([]rune(field), 50, 3, true) {
		t.Errorf("Should be false")
	}

	if shipCouldBePlacedHere([]rune(field), 0, 4, true) {
		t.Errorf("Should be false")
	}

	/*
		_ _ X _ _ X X X X _
		_ _ _ _ _ _ _ _ _ _
		_ _ _ _ _ _ X X _ _
		_ X X _ _ _ _ _ _ X
		_ _ _ _ _ _ _ _ _ _
		_ _ _ _ _ _ _ _ _ _
		_ _ _ X X X _ _ _ _
		X X _ _ _ _ _ _ _ _
		_ _ _ _ _ _ _ _ _ _
		X _ _ _ _ _ X X X _
	*/

	field="OO   OOOO                 OO   OO      O                       OOO    OO                  O     OOO "
	if shipCouldBePlacedHere([]rune(field), 10, 1, true) {
		t.Errorf("Should be false")
	}
}