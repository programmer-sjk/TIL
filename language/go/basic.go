package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	basicGrammer()

	fmt.Println("Bye, World!")
}

func basicGrammer() {
	color := "red"

	if color == "red" {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Color is not red")
	}

	color = "blue"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")
	}

	const name = "JeongKuk"
	fmt.Println(name)

	result := sum(1, 2)
	result = multiply(result, 3)
	fmt.Printf("result = %d\n", result)

	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	result, err := divide(10, 2)
	fmt.Println(result, err)

	result, err = divide(10, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

}

func sum(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}
