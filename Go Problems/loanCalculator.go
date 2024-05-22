/*package main

import "fmt"

func main(){
	var principalAmount, rate float64
	var loanTerm int

	fmt.Println("Enter the principal amount:")
	fmt.Scanln(&principalAmount)

	fmt.Println("Enter the annual interest rate:")
	fmt.Scanln(&rate)

	fmt.Println("Enter the loan term (in months):")
	fmt.Scanln(&loanTerm)

	monthlyPayment:= calculateMonthlyPayment(principalAmount,rate,loanTerm)
	fmt.Println("The monthly payment is:",monthlyPayment)
}

func calculateMonthlyPayment(principalAmount,rate float64, loanTerm int) float64{
	monthlyInterestRate:= (rate/100)/12
	term := float64(loanTerm)

	totalRepayment := principalAmount * (1 + monthlyInterestRate * term)
	monthlyPayment := totalRepayment / term

	return monthlyPayment


}/*