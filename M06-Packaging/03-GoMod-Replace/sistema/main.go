package main

import "github.com/amauryeuzebio/goexpert/M06-Packaging/03-GoMod-Replace/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}
