package main

import "fmt"

func main() {
	p := 74
	q := 47
	fmt.Println("Before:", p, q)
	p, q = q, p

	fmt.Println("After :", p, q)
}
 