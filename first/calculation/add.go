package calculation

// capital name = public function
// untuk package yg berbeda
func Add(number int, numberTwo int) int {
	return add(number, numberTwo)
}

func add(number int, numberTwo int) int {
	return number + numberTwo
}