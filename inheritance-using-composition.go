package main

import (
	//"errors"
	"fmt"
)

type Duck struct {
	ID   int64
	Name string
}

func (d *Duck) Eat() {
	fmt.Println("inside eat function of duck------>", d.Name)
}

type Mallard struct {
	Duck
	Color string
}

func (m *Mallard) Sleep() {
	fmt.Println("inside function of mallard ----->", m.Name)
}

func main() {
	duck := new(Duck)
	duck.ID = 1
	duck.Name = "Pickles"

	mallard := new(Mallard)
	mallard.Color = "Green"
	mallard.Duck = *duck
	mallard.Eat()
    	mallard.Sleep()
}

