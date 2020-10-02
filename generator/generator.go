package generator

import (
	"github.com/mazanax/seabattle/utils"
	"strings"
)

var sizes = [4]uint64{4, 3, 2, 1}

func GenerateField() (string, error) {
	field := []rune(strings.Repeat(string(utils.CellEmpty), 100))
	ships := [4]uint64{1, 2, 3, 4}
	currentShip := 0

	for {
		var pos uint64
		var step uint64
		var vertical bool

		for {
			vertical = false
			step = 1

			probability, err := utils.RandomInt(100)
			if nil != err {
				return "", err
			}

			if probability%2 == 1 {
				vertical = true
				step = 10
			}

			pos, err = utils.RandomInt(100)
			if nil != err {
				return "", err
			}

			if shipCouldBePlacedHere(field, pos, sizes[currentShip], vertical) {
				break
			}
		}

		for i := pos; i < pos+sizes[currentShip]*step; i += step {
			field[i] = utils.CellShip
		}

		ships[currentShip]--
		if ships[currentShip] == 0 {
			currentShip++
		}

		if currentShip >= 4 {
			break
		}
	}

	return string(field), nil
}

func shipCouldBePlacedHere(field []rune, pos uint64, size uint64, vertical bool) bool {
	var startLine uint64
	var endLine uint64
	var step uint64         // шаг, на который надо сместиться, чтобы проверить линию корабля
	var neighborStep uint64 // шаг, на который надо сместиться, чтобы проверить соседние с кораблем клетки

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
	if int(pos-1*step) >= 0 {
		start = pos - 1*step
	}

	for i := start; i < pos+(size+1)*step; i += step {
		if i < startLine || i > endLine {
			continue
		}

		if int(i-neighborStep) > 0 && field[i-neighborStep] != utils.CellEmpty {
			return false
		}

		if field[i] != utils.CellEmpty {
			return false
		}

		if i+neighborStep < 100 && field[i+neighborStep] != utils.CellEmpty {
			return false
		}
	}

	return true
}
