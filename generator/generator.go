package generator

import (
	"github.com/mazanax/seabattle/battlefield"
	"github.com/mazanax/seabattle/utils"
)

var sizes = [4]int8{4, 3, 2, 1}

func GenerateField() ([2]uint64, error) {
	field := battlefield.CreateEmpty()
	ships := [4]int8{1, 2, 3, 4}
	currentShip := 0

	for {
		var pos int8
		var step int8
		var vertical bool

		for {
			vertical = false
			step = 1

			probability, err := utils.RandomInt(100)
			if nil != err {
				return field, err
			}

			if probability%2 == 1 {
				vertical = true
				step = 10
			}

			pos, err = utils.RandomInt8(100)
			if nil != err {
				return field, err
			}

			if shipCouldBePlacedHere(field, pos, sizes[currentShip], vertical) {
				break
			}
		}

		for i := pos; i < pos+sizes[currentShip]*step; i += step {
			field = battlefield.PlaceShip(field, i)
		}

		ships[currentShip]--
		if ships[currentShip] == 0 {
			currentShip++
		}

		if currentShip >= 4 {
			break
		}
	}

	return field, nil
}

func shipCouldBePlacedHere(field [2]uint64, pos int8, size int8, vertical bool) bool {
	var startLine int8
	var endLine int8
	var step int8         // шаг, на который надо сместиться, чтобы проверить линию корабля
	var neighborStep int8 // шаг, на который надо сместиться, чтобы проверить соседние с кораблем клетки

	if !vertical {
		startLine = pos / 10 * 10
		endLine = pos/10*10 + 9
		step = 1
		neighborStep = 10
	} else {
		startLine = pos % 10
		endLine = 90 + pos%10
		step = 10
		neighborStep = 1
	}

	if pos+size*step > endLine {
		return false
	}

	start := pos
	if pos-1*step >= 0 {
		start = pos - 1*step
	}

	for i := start; i < pos+(size+1)*step; i += step {
		if i < startLine || i > endLine {
			continue
		}

		if int(i-neighborStep) > 0 && !battlefield.CellIsEmpty(field, i-neighborStep) {
			return false
		}

		if !battlefield.CellIsEmpty(field, i) {
			return false
		}

		if i+neighborStep < 100 && !battlefield.CellIsEmpty(field, i+neighborStep) {
			return false
		}
	}

	return true
}
