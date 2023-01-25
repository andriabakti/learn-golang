package main

import (
	"fmt"
)

// - parameter/input
// - process
// - return value/output

func base() {
	// luas, _ := calculate(10, 2)
	luas, keliling := calculate2(5, 2)
	fmt.Println(luas)
	fmt.Println(keliling)
}

// multiple return
// func calculate(length, width int) (int, int) {
// 	broad := length * width
// 	girth := 2 * (length + width)

// 	return broad, girth
// }

// predifined return value
func calculate2(length, width int) (broad int, girth int) {
	broad = length * width
	girth = 2 * (length + width)

	return
}