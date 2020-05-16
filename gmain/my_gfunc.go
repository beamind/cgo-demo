package main

import "C"
import "math"

//export MyProd
func MyProd(a int, b int) int {
	return a * b
}

//export MyPow
func MyPow(a float64, b float64) float64 {
	return math.Pow(a, b)
}
