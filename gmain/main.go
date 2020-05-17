package main

// #cgo CFLAGS: -I${SRCDIR}/../include
// #cgo LDFLAGS: -L${SRCDIR}/../lib -lmy_cfunc -lmy_cppfunc
// #include "my_cfunc.h"
import "C"
import "fmt"

func main() {
	x := 12
	y := 23
	z0 := C.my_add(C.int(x), C.int(y))
	z1 := C.my_square(C.int(x))
	fmt.Printf("使用c编译的so库: %d + %d = %d\n", x, y, z0)
	fmt.Printf("使用c++编译的so库: %d^2 = %d\n", x, z1)
}
