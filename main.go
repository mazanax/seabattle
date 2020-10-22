package main

import (
	"fmt"
	"github.com/mazanax/seabattle/generator"
	"github.com/mazanax/seabattle/utils"
)

func renderField(field [2]uint64) {
	for i := 0; i < 100; i++ {
		if !utils.ContainsBit(field[i/64], utils.CellShip<<(63-i%64)) {
			fmt.Print("_ ")
		} else {
			fmt.Print("X ")
		}

		if 0 == (i+1)%10 && i > 0 {
			fmt.Println()
		}
	}
}

func main() {
	var field [2]uint64

	tmpField, err := generator.GenerateField()
	if nil != err {
		panic(err)
	}

	field = tmpField

	renderField(field)
	fmt.Println(field)
}
