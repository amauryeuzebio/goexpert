package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Hello, World")
	tamanho, err := f.Write([]byte("Escrevendo no arquivo"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso!, Tamanho: %d bytes\n", tamanho)

	f.Close()

	// Leitura
	arquivo, err := os.ReadFile("arquivo.txt") // vem tipo bytes

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo)) // Convertendo bytes para string

	// ler arquivo grande usando buffer (lendo aos poucos)
	arquivo2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)

	buffer := make([]byte, 3) // criando buffer de 3 bytes

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	// Removendo um arquivo
	err = os.Remove("arquivo.txt")

	if err != nil {
		panic(err)
	}

}
