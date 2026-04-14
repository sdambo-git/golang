package main

import (
	"fmt"
	"math"
)

var Global int = 1234
var AnotherGlobal = -5678

func main() {
	var j int
	i := Global + AnotherGlobal
	fmt.Println("initial j value:", j)
	fmt.Println("j:", j)
	j = Global
	// math.abs() requires a float64 argument.
	// so we typecast the int to float64, and then back to int.
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d, j=%d k=%2f .\n", Global, i, j, k)
}
