package main

import "fmt"

func main() {

	QuadA(5, 3)
	fmt.Println()
	QuadA(5, 1)
	fmt.Println()
	QuadA(1, 1)
	fmt.Println()
	QuadA(1, 5)

}

func QuadA(x int, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	fmt.Print("o")
	for i := 0; i < x-2; i++ {
		fmt.Print("-")
	}
	if x > 1 {
		fmt.Print("o")
	}
	fmt.Println()

	for j := 0; j < y-2; j++ {
		fmt.Print("|")
		for i := 0; i < x-2; i++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Print("|")
		}
		fmt.Println()
	}

	if y > 1 {
		fmt.Print("o")
		for i := 0; i < x-2; i++ {
			fmt.Print("-")
		}
		if x > 1 {
			fmt.Print("o")
		}
		fmt.Println()
	}
}
