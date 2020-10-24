package validator

import (
	"github.com/mazanax/seabattle/battlefield"
	"github.com/mazanax/seabattle/utils"
	"strings"
	"testing"
)

func TestValidateField(t *testing.T) {
	var field [2]uint64

	field = createFieldWithOnes(15)
	if ValidateField(field) {
		t.Errorf("Field with 15 ones should be invalid")
	}

	field = createFieldFromString(
		"" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ _ _ X X X _ _" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ X _ _ X _ X _" +
			"_ _ _ _ _ _ X _ _ _" +
			"_ _ _ X _ _ X _ X _" +
			"_ _ _ X X _ _ _ _ _" +
			"_ _ _ _ _ _ X X _ _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _",
	)
	if ValidateField(field) {
		t.Errorf("Field contains invalid ships")
	}

	field = createFieldFromString(
		"" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ _ _ X X X _ _" +
			"_ _ _ X _ _ _ _ _ _" +
			"_ X _ X _ _ X _ X _" +
			"_ X _ _ _ _ X _ _ _" +
			"_ _ _ X _ _ X _ X _" +
			"_ _ X X _ _ _ _ _ _" +
			"_ _ _ _ _ _ X X _ _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _",
	)
	if ValidateField(field) {
		t.Errorf("Field doesn't contain all ships")
	}

	field = createFieldFromString(
		"" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ _ _ X X X _ _" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ X _ _ X _ X _" +
			"_ _ _ _ _ _ X _ _ _" +
			"_ _ _ X _ _ X _ X _" +
			"_ _ X X _ _ _ _ _ _" +
			"_ _ _ _ _ _ X X _ _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _",
	)
	if ValidateField(field) {
		t.Errorf("Field contains invalid ships")
	}

	field = createFieldFromString(
		"" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ _ _ X X X _ _" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ X _ _ X _ X _" +
			"_ _ _ _ _ _ X _ _ _" +
			"_ _ _ X _ _ X _ X _" +
			"_ _ X X _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ X _ _" +
			"_ _ _ X _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _",
	)
	if ValidateField(field) {
		t.Errorf("Field contains invalid ships")
	}

	field = createFieldFromString(
		"" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ _ _ X X X _ _" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ X _ _ X _ X _" +
			"_ X _ _ _ _ X _ _ _" +
			"_ _ _ X _ _ X _ X _" +
			"_ _ _ X _ _ _ _ _ _" +
			"_ _ _ _ _ _ X X _ _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _",
	)
	if ValidateField(field) {
		t.Errorf("Field has incorrect ships")
	}

	field = createFieldFromString(
		"" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ _ _ X X X _ _" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ X _ _ X _ X _" +
			"_ _ _ _ _ _ X _ _ _" +
			"_ _ _ X _ _ X _ X _" +
			"_ _ _ X _ _ _ _ _ _" +
			"_ _ X _ _ _ X X _ _" +
			"_ _ _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _",
	)
	if ValidateField(field) {
		t.Errorf("Ships placed too close")
	}

	field = createFieldFromString(
		"" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ _ _ X X X _ _" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ X _ _ X _ X _" +
			"_ _ _ _ _ _ X _ _ _" +
			"_ _ _ X _ _ X _ X _" +
			"_ _ _ X _ _ _ _ _ _" +
			"_ _ _ _ _ _ X X _ _" +
			"_ _ X _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _",
	)
	if !ValidateField(field) {
		t.Errorf("False positive: field is correct")
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

func createFieldWithOnes(onesCount uint8) [2]uint64 {
	result := [2]uint64{0, 0}
	var generatedOnes uint8

	for generatedOnes = 0; generatedOnes < onesCount; generatedOnes++ {
		position, _ := utils.RandomInt8(100)

		result = battlefield.PlaceShip(result, position)
	}

	return result
}
