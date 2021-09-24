package main

import "fmt"

func main() {
	max := 10
	for i := 0; i <= max; i++ {
		for j := 0; j < max-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}
