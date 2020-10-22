package generator

import (
	"github.com/mazanax/seabattle/utils"
	"testing"
)

func TestShipCouldBePlacedHere(t *testing.T) {
	var field []byte

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

	field = []byte{utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty}
	if shipCouldBePlacedHere(field, 2, 1, false) {
		t.Errorf("Should be false")
	}

	if shipCouldBePlacedHere(field, 16, 1, false) {
		t.Errorf("Should be false")
	}

	if shipCouldBePlacedHere(field, 50, 3, true) {
		t.Errorf("Should be false")
	}

	if shipCouldBePlacedHere(field, 0, 4, true) {
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

	field = []byte{utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty}
	if shipCouldBePlacedHere([]byte(field), 10, 1, true) {
		t.Errorf("Should be false")
	}
}
