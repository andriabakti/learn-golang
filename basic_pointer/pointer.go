package main

// pointer => variable yg berisi alamat memori dari suatu nilai
func base() {
	// numberA := 5
	// - referrencing
	// numberB := &numberA

	// println(numberA)
	// println(numberB)
	// - dereferrencing
	// println(*numberB)

	// *numberB = 10
	// println(*numberB)
	// println(numberA)

	// var numberC int = 5
	// var numberD *int = &numberC

	// println(numberC)
	// println(numberD)
	// println(*numberD)

	// numberC = 20
	// println(numberC)
	// println(numberD)
	// println(*numberD)

	number := 6
	println("nilai awal >", number)
	println("alamat memory >", &number)
	changeVal(&number, 13)
	println("nilai baru >", number)
	println("alamat memory >", &number)
}

func changeVal(initVal *int, newVal int) {
	*initVal = newVal
	println("nilai di dalam function >", *initVal)
	println("alamat memory >", initVal)
}