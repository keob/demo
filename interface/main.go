package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func descibe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	i = &T{S: "Hello"}
	descibe(i)
	i.M()

	i = F(math.Pi)
	descibe(i)
	i.M()
}
