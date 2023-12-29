// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello main func")
// 	greet()
// 	_, me := add(1, 2)
// 	fmt.Println(me)
// 	fmt.Println(add(3, 4))
// 	fmt.Println(Proadd(1, 2, 3, 4, 5, 6, 7, 8, 9, 3, 45456, 2463, 6436547, 54757))
// }

// func greet() {
// 	fmt.Println("Good morning")
// }

// func Proadd(V ...int) int {
// 	t := 0

// 	for _, val := range V {
// 		t += val
// 	}

// 	return t
// }

// func add(a int, b int) (int, string) {
// 	return a + b, "helooooooo"
// }

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=-=-

package main

import "fmt"

func main() {

	tanmay := Employee{"Tanmay", "tanmay@gmail.com", false, 24}

	fmt.Println(tanmay)
	fmt.Printf("tanmay details are: %v\n", tanmay)
	tanmay.MethodName()
}

type Employee struct {
	Name string
	EId  string
	Vac  bool
	Age  int
}

func (v Employee) MethodName() {
	fmt.Println("Is User active :", v.Age, v.Vac)
}
