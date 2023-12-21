package main

import "fmt"

func main() {

	m := make(map[int]string) //key: int ... value: string
	m[1] = "tanmay"
	m[3] = "mahato"
	m[5] = "raj"
	m[7] = "aditya"
	fmt.Println("list of values", m)
	fmt.Println(m[3])
	delete(m, 3)
	fmt.Println(m)

	for key, value := range m {
		fmt.Printf("For the key %v, the value given is %v \n", key, value)
	}

	//=-=-=-=-=-=-==-==-=-==-=-=--=-=--=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==--=-=-=-=-=-=-=-=-==--=-=-=
	tanmay := Employee{"Tanmay", "tanmay@gmail.com", false, 24}

	fmt.Println(tanmay)
	fmt.Printf("tanmay details are: %v\n", tanmay)

}

type Employee struct {
	Name string
	EId  string
	Vac  bool
	Age  int
}
