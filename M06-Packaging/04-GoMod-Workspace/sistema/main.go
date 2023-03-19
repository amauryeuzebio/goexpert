package main

import (
	"github.com/amauryeuzebio/goexpert/M06-Packaging/04-GoMod-Workspace/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())

	println(uuid.New().String())

}
