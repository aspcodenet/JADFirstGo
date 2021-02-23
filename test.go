package main

import (
	"fmt"
)

func main() {
	var namn string
	var age int

	bestPlayer := "Mats Sundin"

	fmt.Println("Best player is " + bestPlayer)

	fmt.Println("Vad heter du?")
	fmt.Scanf("%s", &namn)
	fmt.Println("Hur gammal är du?")
	fmt.Scan(&age)

	if age > 48 {
		fmt.Println("Åh vad gammal du är")
	} else {
		fmt.Println("Åh vad ung du är")
	}

	for i := 0; i < age; i++ {
		fmt.Printf("Varv %d\n", i)
	}

}
