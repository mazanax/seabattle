package battlefield

import "github.com/mazanax/seabattle/utils"

const CellEmpty = 0
const CellShip = 1

func CreateEmpty() [2]uint64 {
	return [2]uint64{CellEmpty, CellEmpty}
}

func PlaceShip(field [2]uint64, cell int8) [2]uint64 {
	field[cell/64] |= CellShip << (63 - cell%64)

	return field
}

func CellIsEmpty(field [2]uint64, cell int8) bool {
	return !utils.ContainsBit(field[cell/64], CellShip<<(63-cell%64))
}

func ParseShips(field [2]uint64) [10][]uint8 {
	var ships [10][]uint8
	var shipIterator int8 = 0

	var buffer []uint8
	var i int8

	var offset int8 = 0
	var step int8 = 1

	for i = 0; i < 100; i++ {
		if CellIsEmpty(field, i) {
			continue
		}

		// если в СТРОКЕ есть непустая клетка слева, значит, мы уже проверили текущую клетку
		if i-1 > 0 && (i-1)%10 != 9 && !CellIsEmpty(field, i-1) {
			continue
		}

		// если в СТОЛБЦЕ есть непустая клетка сверху, значит, мы уже проверили текущую клетку
		if i-10 > 0 && !CellIsEmpty(field, i-10) {
			continue
		}

		if (i+1) < 100 && ((i+1)%10 > i%10) && !CellIsEmpty(field, i+1) {
			step = 1
		}

		if (i+10) < 100 && !CellIsEmpty(field, i+10) {
			step = 10
		}

		for {
			if CellIsEmpty(field, i+step*offset) ||
				(step == 1 && i%10 > (i+step*offset)%10) { // дошли до конца линии
				ships[shipIterator] = buffer
				shipIterator++

				buffer = nil
				offset = 0
				step = 1
				break
			}

			buffer = append(buffer, uint8(i+step*offset))
			offset++
		}
	}

	return ships
}
