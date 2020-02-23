package main

import (
	"fmt"

	calculator "../my-package"
)

func main() {
	fmt.Println("Hello Calculator")

	result := calculator.Sum(4, 5)

	fmt.Println("Result of 4 + 5 is ", result)
}
