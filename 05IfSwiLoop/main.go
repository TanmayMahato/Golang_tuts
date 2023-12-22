package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	l := 210
	var r int
	if rt := 4; l < 13 {
		fmt.Println("hehe")
		fmt.Println(rt)
		r = 3
	} else if l > 13 {
		fmt.Println("noo")
	} else {
		fmt.Println("noo")
		r = 5
	}
	fmt.Println(r)

	//================================

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3) + 1

	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}

}
