package main

import "fmt"

func main() {
	max := 10
	for i := 1; i <= max; i++ {
		var nLeft int
		if max-i%2 == 0 {
			nLeft = (max - i) / 2

		} else {
			nLeft = (max - i) / 2
		}
		nRight := max - nLeft - i
		for j := 1; j <= max; j++ {
			if j <= nLeft || j >= max-nRight {
				fmt.Print("  ")
			} else {
				fmt.Print(" * ")
			}
		}
		fmt.Println()
	}
}
