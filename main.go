package main

import (
	"fmt"
	"github.com/mazanax/seabattle/battlefield"
	"github.com/mazanax/seabattle/generator"
	"github.com/mazanax/seabattle/validator"
)

func renderField(field [2]uint64) {
	var i int8
	for i = 0; i < 100; i++ {
		if battlefield.CellIsEmpty(field, i) {
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
	fmt.Println(validator.ValidateField(field))
}
