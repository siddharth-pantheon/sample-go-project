package test

import (
	"fmt"
	"sample-go/utils"
	"testing"

	"github.com/dailymotion/allure-go"
)

func TestCalculator(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		allure.Step(allure.Description("Simple Calculator"),
			allure.Action(func() {

				a := 10
				b := 2

				sum := utils.Addition(a, b)
				fmt.Print("\nSummation is: ", sum)

				diff := utils.Substraction(a, b)
				fmt.Print("\nDifference is: ", diff)

				mul := utils.Multiplication(a, b)
				fmt.Print("\nMultiplication is: ", mul)

				div := utils.Division(a, b)
				fmt.Print("\n Division is: ", div)

				y, z := utils.ReturnIntVals()
				fmt.Println(y)
				fmt.Println(z)

				s := utils.ReturnStringVals()
				fmt.Print(s)

			}))
	}))
}

func TestFindLargestNumber(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		allure.Step(allure.Description("Finding largest number"),
			allure.Action(func() {
				first := 15
				second := 20
				third := 8

				largestNumber := utils.FindLargestValue(first, second, third)
				fmt.Println("Largest among three numbers is :", largestNumber)

			}))
	}))
}
