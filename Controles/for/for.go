package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // Não existe while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { // i += incrmenta de 2 em 2
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Ímpar ")
		}
	}

	for {
		// laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
