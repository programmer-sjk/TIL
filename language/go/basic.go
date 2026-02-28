package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("===== basic grammar ======")
	basicGrammer()

	fmt.Printf("\n")
	
	fmt.Println("===== collections ======")
	collections()
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

	for i := 0; i < 3; i++ {
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

type User struct {
	name string
	age  int
}

func (u User) greeting() string {
	return "Greeting"
}


func collections() {
	arr := []int{1, 2, 3}
	arr = append(arr, 4)
	fmt.Println(arr, arr[0])

	scores := map[string]int{
		"정국": 100,
		"영희": 90,
		"철수": 80,
	}

	fmt.Println(scores, scores["정국"])

	user := User{name: "정국", age: 20}
	fmt.Println(user, user.name, user.age, user.greeting())
	changeAge(&user)
	fmt.Println(user, user.name, user.age, user.greeting())

}

func changeAge(user *User) {
	user.age = 21
}

\