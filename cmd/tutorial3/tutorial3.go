package main

import (
	"errors"
	"fmt"
)

func main() {

	text := "Hey  wassup"
	printString(text)

	num1 := 34
	num2 := 535

	var result, remainder = intDivision(num1, num2)
	fmt.Println(result, remainder)

	fmt.Println("Now calling the function with error checks !")

	num3, num4 := 434, 0
	result, remainder, err := intDivErrorchk(num3, num4)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The result is %v and the remainder is %v .", result, remainder)

	}

}

func printString(text string) {
	fmt.Println(text)
}

func intDivision(num1 int, num2 int) (int, int) {
	result := num1 / num2
	remainder := num1 % num2

	fmt.Printf("The result of the operation is %v and the remainder for the same is %v .", result, remainder)

	return result, remainder
}

func intDivErrorchk(num1 int, num2 int) (int, int, error) {
	var err error
	if num2 == 0 {
		err = errors.New(" The Demnominator Cant Be Zero! ")
		return 0, 0, err
	}

	return num1 / num2, num1 % num2, err
}
