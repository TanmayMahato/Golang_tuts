//DEFER  ..DEFER

package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("\ntwo")
	fmt.Println("Hello")
	rev()
}

func rev() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
