package main

import "fmt"

type opFuncType func(int, int) int

func add(i int, j int) int { return i + j }

func sub(i int, j int) int { return i - j }

func mul(i int, j int) int { return i * j }

func div(i int, j int) int { return i / j }

func main() {
	opMap := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	opMapw := map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	fmt.Println(opMap["+"](4, 5))
	fmt.Println(opMap["*"](4, 5))

	fmt.Println(opMapw["+"](12, 12))
	fmt.Println(opMapw["*"](12, 12))

}
