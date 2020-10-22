package main

import (
	"fmt"
	"github.com/mazanax/seabattle/generator"
	"github.com/mazanax/seabattle/utils"
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
	var field []byte

	tmpField, err := generator.GenerateField()
	if nil != err {
		panic(err)
	}

	field = tmpField

	renderField(string(field))
}
