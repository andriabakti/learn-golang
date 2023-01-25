package basic_map_slice

import "fmt"

func sliceOfMap() {
	students := []map[string]string{
		{"name": "Lecca", "score": "A"},
		{"name": "Myna", "score": "C"},
	}

	for _, val := range students {
		fmt.Println(val["name"])
	}
}