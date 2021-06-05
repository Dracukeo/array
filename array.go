package main

import "fmt"

func main() {
	var a [5]int
	fmt.Print("a")

	a[2] = 7
	fmt.Print("a")

	a := [5]int{5, 4, 3, 2, 1}
	fmt.Print("a")

	s := []int{5, 4, 3, 2, 1}
	fmt.Print("s")

	s = append(s, 13)
	fmt.Print("s")
}
