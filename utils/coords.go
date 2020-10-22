package utils

import "strconv"

var letterToDigit = map[byte]int{
	'A': 0,
	'B': 1,
	'C': 2,
	'D': 3,
	'E': 4,
	'F': 5,
	'G': 6,
	'H': 7,
	'I': 8,
	'J': 9,
}
var digitToLetter = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

func ToHumanReadable(coord int) string {
	return digitToLetter[coord%10] + strconv.Itoa(coord/10+1)
}

func FromHumanReadable(coord string) int {
	y, _ := strconv.Atoi(coord[1:])

	return (y-1)*10 + letterToDigit[coord[0]]
}
