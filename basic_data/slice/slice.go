package basic_slice

import "fmt"

func testSlice() {
	// slice
	var devices []string
	// devices := []string{"Android", "iPhone"}
	devices = append(devices, "iPad")
	devices = append(devices, "PC")

	for _, item := range devices {
		fmt.Println(item)
	}
	// fmt.Println(devices)
}