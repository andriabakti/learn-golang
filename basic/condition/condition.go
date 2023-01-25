package basic_condition

import "fmt"

func testCondition() {
	// if-else
	// age := 20

	// if age > 10 {
	// 	fmt.Println("Boleh main")
	// } else {
	// 	fmt.Println("Gak boleh")
	// }

	score := 80
	var grade string

	if (score <= 50) {
		grade = "D"
	} else if (score <= 60) {
		grade = "C"
	} else if (score < 70) {
		grade = "B"
	} else {
		grade = "A"
	}

	fmt.Println(grade)

	// switch-case
	number := 2

	switch number {
	case 1:
		fmt.Println("Hi")
	case 2:
		fmt.Println("Fu")
	case 3:
		fmt.Println("Mi")
	default:
		fmt.Println("Hoka")
	}

	// if number == 1 {
	// 	fmt.Println("Il")
	// } else if number == 2 {
	// 	//
	// } else if number == 3 {
	// 	//
	// }
}