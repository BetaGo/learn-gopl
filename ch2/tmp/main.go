package main

import "fmt"

// func main() {
// 	x := 1
// 	p := &x
// 	fmt.Println(*p)
// 	*p = 2
// 	fmt.Println(x)
// }

func main() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)
}
