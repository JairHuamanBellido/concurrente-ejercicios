package main

import "fmt"

func shadowTutorial() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 9
		fmt.Println(x)
	}
	fmt.Println(x)
}
func main() {
	shadowTutorial()
}
