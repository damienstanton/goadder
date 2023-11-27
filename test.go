package main

import (
	"fmt"

	"code.pwc.com/damienstanton/goadder/src/uniffi/math"
)

func main() {
	four := math.Add(2, 2)
	fmt.Printf("2 + 2 : %d\n", four)
}
