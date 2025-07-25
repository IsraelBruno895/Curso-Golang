package main

import "fmt"

func main() {
	var a int     // 0
	var b float64 // 0
	var c bool    // false
	var d string  // "" string vazia
	var e *int    // nil na maioria das linguagens é null mas em GO é <nil>

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)

}
