package main

import "fmt"

func main() {
	var temp float64
	

	fmt.Println("Digite a temperatura: ")
	fmt.Scan(&temp)

	if temp > 30{
		fmt.Printf("Esta muito quente\n")
	} else if temp > 20 && temp <= 30 {
		fmt.Printf("Temperatura agradavel")
	} else {
		fmt.Printf("Esta muito frio")
	}
	
	


}
