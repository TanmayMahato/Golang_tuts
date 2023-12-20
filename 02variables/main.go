package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var usrnm string = "tanmay"
	// fmt.Println(usrnm)
	// fmt.Printf("variable is of type: %T \n", usrnm)

	// var t float32 = 23.231321232321313123
	// fmt.Println(t)
	// fmt.Printf("variable is of type : %T \n", t)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the string :")

	input, _ := reader.ReadString('.')
	fmt.Println(input, "hehe")

	var myPointer *int
	fmt.Println("value", myPointer)

	pt := 23
	var ptr = &pt
	fmt.Println("value", ptr)
}

// go help and go env
//...go build...
// GOOS ="windows" go build          creates a windows executable file

// Pointers
