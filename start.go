package main

import (
	"fmt"
	"github.com/stevemasta34/neuralnetwork/util"
)

func main() {
	fmt.Println("Just making sure things compile")
	left, right := util.TestFunc("car")
	fmt.Printf("%v\n", left)
	fmt.Printf("%v\n", right)
}
