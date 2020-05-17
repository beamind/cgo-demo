package main

import "C"
import (
	"math"
	"unicode/utf8"
)

//export MyProd
func MyProd(a int, b int) int {
	return a * b
}

//export MyPow
func MyPow(a float64, b float64) float64 {
	return math.Pow(a, b)
}

//export MyStrLen
func MyStrLen(s *C.char) int {
	return utf8.RuneCountInString(C.GoString(s))
}
