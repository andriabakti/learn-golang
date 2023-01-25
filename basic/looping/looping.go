package main

import "fmt"

func main() {
	// - for i
	// for i := 1; i <= 20; i++ {
	// 	fmt.Println("Saya sedang belajar Go :", i)
	// }

	// - for i (seperti while)
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("Saya sedang belajar Go :", i)
	// 	i++
	// }

	// - for range (seperti forEach)
	//  untuk array
	title := "Golang the best language"
	// for i, letter := range title {
	// 	// return range = byte
	// 	if (i % 2 == 0) {
	// 		fmt.Println("Letter :", string(letter))
	// 	}
	// }

	for index, letter := range title {
		letterStr := string(letter)
		// if letterStr == "a" || letterStr == "i" || letterStr == "u" || letterStr == "e" || letterStr == "o" {
		// 	fmt.Println("Index: ", index , " Letter :", string(letter))
		// }
		switch letterStr {
		case "a", "i", "u", "e", "o":
			fmt.Println("Index: ", index , " Letter :", string(letter))
		}
	}
}