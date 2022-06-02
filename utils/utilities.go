package utils

import (
	"fmt"
)

func Addition(a int, b int) int {
	sum := a+b
	return sum
}

func Substraction(a int, b int) int {
	diff := a - b
	return diff
}

func Multiplication(a int, b int) int {
	mul := a * b
	return mul
}

func Division(a int, b int) int {
	div := a / b
	return div
}

func ReturnIntVals() (int, int) {
	value1, value2 := 3,7
	return value1, value2
}

func ReturnStringVals() string {
	sampleString := "Hello World"
	return sampleString
}

type Rect struct {
	width, height int
}

func (r *Rect) Area() int {
	area := r.width * r.height
	return area
}

func (r Rect) Perim() int {
	perim:= 2*r.width + 2*r.height
	return perim
}

func FindLargestValue(first int, second int, third int) (int) {
	/* check the boolean condition using if statement */
	if first >= second && first >= third {
		return first /* if condition is true then print the following */
	}
	if second >= first && second >= third {
		return second
	}
	if third >= first && third >= second {
		return third
	}
	return 0
}

func HelloWorld() (string){
	h := "Hello World"
	return h
}

func GetUrl() (string){
	user := "root"
	password:= "supersecret" // Sensitive
  
	url := "login=" + user + "&passwd=" + password
	return url
  }

  func IncorrectIfCondition() {
	  a:= "test"
	if a == "test" {
		fmt.Print("test")
	} else if a == "test" { 
		fmt.Print("test")
	}
	a=a
	// TO DO
	// FIXME
  }

  func EmptyFunction() {}