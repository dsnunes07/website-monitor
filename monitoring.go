package main

import "fmt"

func main() {
	fmt.Println("1. Iniciar monitoramento")
	fmt.Println("2. Exibir logs")

	var option int
	// fmt.Scanf("%d", &option)
	fmt.Scan(&option)

	if option == 1 {
		fmt.Println("Starting monitoring...")
	} else if option == 2 {
		fmt.Println("Current logs:")
	}
}
