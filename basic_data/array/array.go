package basic_array

import "fmt"

func testArray() {
	// Array
	// - bentuk 1
	// var langs [5]string
	// langs[0] = "TypeScript"
	// langs[1] = "Ruby"
	// langs[2] = "C#"
	// langs[3] = "Python"
	// langs[4] = "Go"

	// - bentuk 2
	langs := [...]string{
		"Ruby",
		"Python",
		"Go",
		"C#",
		"Kotlin",
		"Java",
	}
	fmt.Println(langs)
	length := len(langs)
	fmt.Println(length)

	for _, item := range langs {
		fmt.Println("Item >", item)
	}
}