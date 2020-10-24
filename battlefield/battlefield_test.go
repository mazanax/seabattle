package battlefield

import (
	"reflect"
	"strings"
	"testing"
)

/*
	0			1			2			3			4			5			6			7			8			9
0	A1: 0, 63	B1: 0, 62	C1: 0, 61	D1: 0, 60	E1: 0, 59	F1: 0, 58	G1: 0, 57	H1: 0, 56	I1: 0, 55	J1: 0, 54
1	A2: 0, 53	B2: 0, 52	C2: 0, 51	D2: 0, 50	E2: 0, 49	F2: 0, 48	G2: 0, 47	H2: 0, 46	I2: 0, 45	J2: 0, 44
2	A3: 0, 43	B3: 0, 42	C3: 0, 41	D3: 0, 40	E3: 0, 39	F3: 0, 38	G3: 0, 37	H3: 0, 36	I3: 0, 35	J3: 0, 34
3	A4: 0, 33	B4: 0, 32	C4: 0, 31	D4: 0, 30	E4: 0, 29	F4: 0, 28	G4: 0, 27	H4: 0, 26	I4: 0, 25	J4: 0, 24
4	A5: 0, 23	B5: 0, 22	C5: 0, 21	D5: 0, 20	E5: 0, 19	F5: 0, 18	G5: 0, 17	H5: 0, 16	I5: 0, 15	J5: 0, 14
5	A6: 0, 13	B6: 0, 12	C6: 0, 11	D6: 0, 10	E6: 0, 9	F6: 0, 8	G6: 0, 7	H6: 0, 6	I6: 0, 5	J6: 0, 4
6	A7: 0, 3	B7: 0, 2	C7: 0, 1	D7: 0, 0	E7: 1, 63	F7: 1, 62	G7: 1, 61	H7: 1, 60	I7: 1, 59	J7: 1, 58
7	A8: 1, 57	B8: 1, 56	C8: 1, 55	D8: 1, 54	E8: 1, 53	F8: 1, 52	G8: 1, 51	H8: 1, 50	I8: 1, 49	J8: 1, 48
8	A9: 1, 47	B9: 1, 46	C9: 1, 45	D9: 1, 44	E9: 1, 43	F9: 1, 42	G9: 1, 41	H9: 1, 40	I9: 1, 39	J9: 1, 38
9	A10: 1, 37	B10: 1, 36	C10: 1, 35	D10: 1, 34	E10: 1, 33	F10: 1, 32	G10: 1, 31	H10: 1, 30	I10: 1, 29	J10: 1, 28
*/

func TestCellIsShip(t *testing.T) {
	field := [2]uint64{1 << 63, 1 << 50}

	if !CellIsEmpty(field, 63) {
		t.Errorf("Cell 63 is empty")
	}

	if CellIsEmpty(field, 0) {
		t.Errorf("Cell 0 is empty")
	}

	if CellIsEmpty(field, 77) {
		t.Errorf("Cell 77 is not empty")
	}

	if !CellIsEmpty(field, 76) {
		t.Errorf("Cell 76 is empty")
	}
}

func TestParseShips(t *testing.T) {
	var field [2]uint64
	var ships [10][]uint8

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
	ships = ParseShips(field)

	if !reflect.DeepEqual(ships[0], []uint8{1, 11, 21, 31}) {
		t.Errorf("The 1st ship placed in [1, 11, 21, 31], but got %+v", ships[0])
	}

	if !reflect.DeepEqual(ships[1], []uint8{3}) {
		t.Errorf("The 2nd ship placed in [3], but got %+v", ships[1])
	}

	if !reflect.DeepEqual(ships[2], []uint8{15, 16, 17}) {
		t.Errorf("The 3rd ship placed in [15, 16, 17], but got %+v", ships[2])
	}

	if !reflect.DeepEqual(ships[3], []uint8{23, 33}) {
		t.Errorf("The 4th ship placed in [23, 33], but got %+v", ships[3])
	}

	if !reflect.DeepEqual(ships[4], []uint8{36, 46, 56}) {
		t.Errorf("The 5th ship placed in [36, 46, 56], but got %+v", ships[4])
	}

	if !reflect.DeepEqual(ships[5], []uint8{38}) {
		t.Errorf("The 5th ship placed in [38], but got %+v", ships[5])
	}

	if !reflect.DeepEqual(ships[6], []uint8{53, 63}) {
		t.Errorf("The 6th ship placed in [53, 63], but got %+v", ships[6])
	}

	if !reflect.DeepEqual(ships[7], []uint8{58}) {
		t.Errorf("The 7th ship placed in [58], but got %+v", ships[7])
	}

	if !reflect.DeepEqual(ships[8], []uint8{76, 77}) {
		t.Errorf("The 8th ship placed in [76, 77], but got %+v", ships[8])
	}

	if !reflect.DeepEqual(ships[9], []uint8{82}) {
		t.Errorf("The 9th ship placed in [82], but got %+v", ships[9])
	}

	field = createFieldFromString(
		"" +
			"_ X _ _ _ _ _ _ _ _" +
			"_ X _ _ _ X X X _ _" +
			"_ X _ X _ _ _ _ _ _" +
			"_ X _ X _ _ X _ X _" +
			"_ _ _ _ _ _ X _ _ _" +
			"_ _ _ X _ _ X _ X _" +
			"_ _ _ X _ _ _ _ _ _" +
			"_ _ _ _ _ _ X X _ X" +
			"X _ _ _ _ _ _ _ _ _" +
			"_ _ _ _ _ _ _ _ _ _",
	)
	ships = ParseShips(field)

	if !reflect.DeepEqual(ships[0], []uint8{1, 11, 21, 31}) {
		t.Errorf("The 1st ship placed in [1, 11, 21, 31], but got %+v", ships[0])
	}

	if !reflect.DeepEqual(ships[1], []uint8{15, 16, 17}) {
		t.Errorf("The 2nd ship placed in [15, 16, 17], but got %+v", ships[1])
	}

	if !reflect.DeepEqual(ships[2], []uint8{23, 33}) {
		t.Errorf("The 3rd ship placed in [23, 33], but got %+v", ships[3])
	}

	if !reflect.DeepEqual(ships[3], []uint8{36, 46, 56}) {
		t.Errorf("The 4th ship placed in [36, 46, 56], but got %+v", ships[4])
	}

	if !reflect.DeepEqual(ships[4], []uint8{38}) {
		t.Errorf("The 5th ship placed in [38], but got %+v", ships[4])
	}

	if !reflect.DeepEqual(ships[5], []uint8{53, 63}) {
		t.Errorf("The 6th ship placed in [53, 63], but got %+v", ships[5])
	}

	if !reflect.DeepEqual(ships[6], []uint8{58}) {
		t.Errorf("The 7th ship placed in [58], but got %+v", ships[6])
	}

	if !reflect.DeepEqual(ships[7], []uint8{76, 77}) {
		t.Errorf("The 8th ship placed in [76, 77], but got %+v", ships[7])
	}

	if !reflect.DeepEqual(ships[8], []uint8{79}) {
		t.Errorf("The 9th ship placed in [79], but got %+v", ships[8])
	}

	if !reflect.DeepEqual(ships[9], []uint8{80}) {
		t.Errorf("The 10th ship placed in [80], but got %+v", ships[9])
	}
}

func createFieldFromString(fieldStr string) [2]uint64 {
	var field [2]uint64

	fieldStr = strings.ReplaceAll(fieldStr, " ", "")

	for i := 0; i < 100; i++ {
		if fieldStr[i] == 'X' {
			field = PlaceShip(field, int8(i))
		}
	}

	return field
}
