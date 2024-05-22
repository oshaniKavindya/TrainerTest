package main

import (
	"fmt"
)

func main(){
	var num1,num2 float64
	var operator string

	fmt.Println("Enter number1:")
	fmt.Scanln(&num1)

	fmt.Println("Enter number 2:")
	fmt.Scanln(&num2)

	fmt.Println("Enter the operator [+ , - , / , *] :")
	fmt.Scanln(&operator)

	result := calculator(num1, num2, operator)
	fmt.Println("Result is :", result)

	
}
func calculator(num1, num2 float64, operator string) float64 {
		var result float64

		if operator =="+"{
			result = num1+num2
		} else if operator == "-" {
			result =num1-num2
		} else if operator == "/" {
			if num2 != 0 {
				result = num1 / num2
			} else {
				fmt.Println("Error: Division by zero")
				return 0
			}
		}else if operator =="*"{
			result= num1*num2
		}
		return result
	}
