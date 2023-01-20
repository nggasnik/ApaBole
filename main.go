package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {

	if i%3 == 0 {
		fmt.Printf("Apa")
	}
	if i%5 == 0 {
		fmt.Printf("Bole")
	}
	if i%3 != 0 && i%5 !=0 {
		fmt.Printf("%d", i)
	}
		fmt.Printf("\n")
	}
}