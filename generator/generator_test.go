package generator

import (
	"github.com/mazanax/seabattle/utils"
	"testing"
)

func TestShipCouldBePlacedHere(t *testing.T) {
	field := [2]uint64{utils.CellEmpty, utils.CellEmpty}
	var fieldTemp []byte

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

	fieldTemp = []byte{utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty}
	for i := 0; i < 100; i++ {
		if fieldTemp[i] == utils.CellEmpty {
			continue
		}

		field[i/64] |= 1 << (63 - i%64)
	}

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

	fieldTemp = []byte{utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellEmpty, utils.CellShip, utils.CellShip, utils.CellShip, utils.CellEmpty}
	for i := 0; i < 100; i++ {
		if fieldTemp[i] == utils.CellEmpty {
			continue
		}

		field[i/64] |= 1 << (63 - i%64)
	}

	if shipCouldBePlacedHere(field, 10, 1, true) {
		t.Errorf("Should be false")
	}
}
