package main

import (
	"fmt"
	"seabattle/generator"
	"seabattle/utils"
)

func renderField(field string) {
	for i := 0; i < 100; i++ {
		if utils.CellEmpty == field[i] {
			fmt.Print("_ ")
		}

		if utils.CellShip == field[i] {
			fmt.Print("X ")
		}

		if 0 == (i+1)%10 && i > 0 {
			fmt.Println()
		}
	}
}

func main() {
	var field [2]string

	tmpField, err := generator.GenerateField()
	if nil != err {
		panic(err)
	}

	field[0] = tmpField

	tmpField, err = generator.GenerateField()
	if nil != err {
		panic(err)
	}

	field[1] = tmpField

	renderField(field[0])
	fmt.Println()
	fmt.Println()
	renderField(field[1])
}
