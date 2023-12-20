package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)

	var arr1 = [3]int{10, 11, 12}
	fmt.Println(arr1)

	var sli = []int{22, 21, 23}

	fmt.Println(sli)

	var sl1 []int
	sl1 = append(sl1, 31, 32, 33, 34)
	fmt.Println(sl1)

	sl1 = append(sl1[1:])
	fmt.Println(sl1)

	hs := make([]int, 4)
	hs[0] = 91
	hs[1] = 94
	hs[2] = 92
	hs[3] = 90

	hs = append(hs, 98, 100)
	fmt.Println(hs)

	sort.Ints(hs)
	fmt.Println(hs)
	fmt.Println(sort.IntsAreSorted(hs))

	//removing....

	var a = []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println(a)

	var i int = 3
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)

}
