package generator

import (
	"github.com/mazanax/seabattle/battlefield"
	"math/bits"
	"strings"
	"testing"
)

func TestShipCouldBePlacedHere(t *testing.T) {
	field := battlefield.CreateEmpty()

	field = createFieldFromString(
		"" +
			"_ _ X _ _ X X X X _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ X X _ _" +
			"_ X X _ _ _ _ _ _ X" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ X X X _ _ _ _" +
			"X X _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"X _ _ _ _ _ X X X _",
	)

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

	field = createFieldFromString(
		"" +
			"X X X _ _ X X X X _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ X X _ _" +
			"_ X X _ _ _ _ _ _ X" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ X X X _ _ _ _" +
			"X X _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"X _ _ _ _ _ X X X _",
	)

	if shipCouldBePlacedHere(field, 10, 1, true) {
		t.Errorf("Should be false")
	}
}

func TestCountOfOnes(t *testing.T) {
	field, _ := GenerateField()
	cellsCount := bits.OnesCount64(field[0]) + bits.OnesCount64(field[1])

	if cellsCount != 20 {
		t.Errorf("Number of `Ship` cells should be 20. But %d generated.", cellsCount)
	}
}

func createFieldFromString(fieldStr string) [2]uint64 {
	var field [2]uint64

	fieldStr = strings.ReplaceAll(fieldStr, " ", "")

	for i := 0; i < 100; i++ {
		if fieldStr[i] == 'X' {
			field = battlefield.PlaceShip(field, int8(i))
		}
	}

	return field
}
