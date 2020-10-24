package validator

import (
	bf "github.com/mazanax/seabattle/battlefield"
	"math/bits"
)

func ValidateField(field [2]uint64) bool {
	if bits.OnesCount64(field[0])+bits.OnesCount64(field[1]) != 20 {
		return false
	}

	if countShips(field) != 10 {
		return false
	}

	ships := bf.ParseShips(field)
	var sizes [4]uint8

	for i := 0; i < 10; i++ {
		if len(ships[i]) == 0 || len(ships[i]) > 4 {
			return false
		}

		// нужно проверить только углы каждой клетки, т.к. остальное покрывается другими условиями
		for j := 0; j < len(ships[i]); j++ {
			corners := [4]int8{
				int8(ships[i][j]) - 11, // -10 - 1 вверхний левый
				int8(ships[i][j]) - 9,  // -10 + 1 верхний правый
				int8(ships[i][j]) + 11, // +10 + 1 нижний правый
				int8(ships[i][j]) + 9,  // +10 - 1 нижний левый
			}

			if corners[0] > 0 && corners[0]%10 != 9 && !bf.CellIsEmpty(field, corners[0]) ||
				corners[1] < 100 && corners[1]/10 == int8(ships[i][j])/10-1 && !bf.CellIsEmpty(field, corners[1]) ||
				corners[2] < 100 && corners[2]/10 == int8(ships[i][j])/10+1 && !bf.CellIsEmpty(field, corners[2]) ||
				corners[3] > 0 && corners[3]%10 != 9 && !bf.CellIsEmpty(field, corners[3]) {
				return false
			}
		}

		sizes[len(ships[i])-1]++
	}

	if sizes[0] != 4 || sizes[1] != 3 || sizes[2] != 2 || sizes[3] != 1 {
		return false
	}

	return true
}

func countShips(field [2]uint64) uint8 {
	var count uint8 = 0
	var i int8

	for i = 0; i < 100; i++ {
		// проверяем, что предыдущая клетка в этой СТРОКЕ не является частью корабля
		if i-1 > 0 && (i-1)%10 != 9 && !bf.CellIsEmpty(field, i-1) {
			continue
		}

		// проверяем, что предыдущая клетка в этом СТОЛБЦЕ не является частью корабля
		if i-10 > 0 && !bf.CellIsEmpty(field, i-10) {
			continue
		}

		if !bf.CellIsEmpty(field, i) {
			count++
		}
	}

	return count
}
