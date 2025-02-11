package main

import "fmt"

func main() {
	var num int
	

	fmt.Println("Digite um numero: ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		fmt.Print("\nContagem: ", i)
	}
	
	


}
