package main

// 2 type of package in Go
// - executable (main)
// - library (selain main)

import (
	"first/calculation"
	"fmt"
)

// fungsi import:
// 1. mengakses standart library go
// 2. mengakses package yg berbeda
// 3. mengakses library yg dibuat orang lain

func main() {
	result := calculation.Add(8, 9)
	fmt.Println("Halo, belajar Golang")
	fmt.Printf("Result: %d", result)
}
// function yg wajib ada untuk mengeksekusi package main