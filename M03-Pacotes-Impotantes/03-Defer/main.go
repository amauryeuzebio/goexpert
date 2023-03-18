package main

import "fmt"

func main() {
	defer fmt.Println("Primeira linha") // defer executa somente no final
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
}
