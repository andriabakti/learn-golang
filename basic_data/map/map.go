package basic_map

import "fmt"

func testMap() {
	// var myMap = map[string]int
	// myMap = map[string]int{}
	// -
	// myMap := map[string]int{}

	// myMap["Ruby"] = 9
	// myMap["Python"] = 3
	// myMap["Ruby"] = 11

	// fmt.Println(myMap)
	// fmt.Println(myMap["Ruby"])

	// -
	myMap2 := map[string]string{
		"Ruby": "is beautiful",
		"Go": "is super fast",
		"Java": "is powerful",
	}

	value, isAvailable := myMap2["python"]
	fmt.Println("value >", value)
	fmt.Println("isAvailable >", isAvailable)
	delete(myMap2, "Ruby")
	for key, val := range myMap2 {
		fmt.Println("Key: ", key, " & value: ", val)
	}
}