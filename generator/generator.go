package generator

import (
	"bytes"
	"github.com/mazanax/seabattle/utils"
)

var sizes = [4]int64{4, 3, 2, 1}

func GenerateField() ([]byte, error) {
	field := bytes.Repeat([]byte{utils.CellEmpty}, 100)
	ships := [4]int64{1, 2, 3, 4}
	currentShip := 0

	for {
		var pos int64
		var step int64
		var vertical bool

		for {
			vertical = false
			step = 1

			probability, err := utils.RandomInt(100)
			if nil != err {
				return nil, err
			}

			if probability%2 == 1 {
				vertical = true
				step = 10
			}

			pos, err = utils.RandomInt(100)
			if nil != err {
				return nil, err
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

	return field, nil
}

func shipCouldBePlacedHere(field []byte, pos int64, size int64, vertical bool) bool {
	var startLine int64
	var endLine int64
	var step int64         // шаг, на который надо сместиться, чтобы проверить линию корабля
	var neighborStep int64 // шаг, на который надо сместиться, чтобы проверить соседние с кораблем клетки

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
