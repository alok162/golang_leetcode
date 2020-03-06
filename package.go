package main

import "fmt"

type Calculator interface {
	addition() int
	multiplication() int
}

type Input struct {
	x, y int
}

func (input Input) addition() int {
	return input.x + input.y
}

func (input Input) multiplication() int {
	return input.x * input.y
}

func getAns(calc Calculator) {
	fmt.Println(calc.addition())
	fmt.Println(calc.multiplication())
} 

func main() {
	input := Input{x:5, y:5}
	getAns(input)
}
