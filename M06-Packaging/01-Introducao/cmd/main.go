package main

import (
	"fmt"

	"github.com/amauryeuzebio/goexpert/M06-Packaging/01-Introducao/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.C)
	// fmt.Println(m.Add())
	// fmt.Println(math.X)
}
