/*package main

import (
	"fmt"
	"strings"
)

func main(){
	var  temp float64
	var unit string

	fmt.Println("Enter the temperature value: ")
	fmt.Scanln(&temp)

	fmt.Println("Enter the unit (C for Celsius, F for Fahrenheit):")
	fmt.Scanln(&unit)
	
	upperCaseUnit := strings.ToUpper(unit)

	if upperCaseUnit =="C"{
		convertTemp:=celToFar(temp)
		fmt.Printf("%.2f°C is %.2f°F \n ",temp,convertTemp)
	}else if upperCaseUnit=="F"{
		convertTemp:=farToCel(temp)
		fmt.Printf("%.2f°F is %.2f°C \n",temp,convertTemp)
	}else{
		fmt.Println("please enter correct unit type")
	}
}

func celToFar (temp float64)float64{
 return (temp*9/5)+32
}

func farToCel (temp float64) float64{
	return (temp-32)* 5/9
}*/