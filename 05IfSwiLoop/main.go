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

	//==================================

	d := []string{"tanmay", "mahato", "raj", "rahul"}

	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}

	//for array and slice

	for j := range d {
		fmt.Println(d[j])
	}

	for index, value := range d { // while using range
		fmt.Printf("index is %v and it's value is %v \n", index, value)
	}

	for q := 0; q < 10; q++ {
		if q == 2 {
			continue
		}
		if q == 4 {

		}

		if q == 6 {
			goto lv
		}
		fmt.Println(q)
	}

lv:
	fmt.Println("you literally jumbper heh")
}
