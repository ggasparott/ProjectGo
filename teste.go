package main

import "fmt"

func main() {
	var idade int
	var nome string
	var cidade string

	fmt.Println("Digite seu nome: ")
	fmt.Scan(&nome)

	fmt.Println("Digite sua cidade: ")
	fmt.Scan(&cidade)

	fmt.Println("Digite sua idade: ")
	fmt.Scan(&idade)

	fmt.Printf("Ola, %s! Voce tem %d anos e mora em %s.\n", nome, idade, cidade)

}
