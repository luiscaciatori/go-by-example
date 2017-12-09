package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	res := sum(vals())
	fmt.Println(res)

}
