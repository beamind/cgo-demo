package main

// #cgo CFLAGS: -I${SRCDIR}/../include
// #cgo LDFLAGS: -L${SRCDIR}/../lib -lmy_cfunc
// #include "my_cfunc.h"
import "C"
import "fmt"

func main() {
	x := 12
	y := 23
	z := C.my_add(C.int(x), C.int(y))
	fmt.Printf("%d + %d = %d\n", x, y, z)
}
