package main

import (
	"fmt"

	"github.com/Exitialis/go-package-test/calculator"
)

func main() {
	fmt.Println(calculator.Add(2, 3))
	fmt.Println(calculator.Sub(2, 3))
	fmt.Println(calculator.Sub(3, 2))
}
