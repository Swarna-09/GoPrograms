//Problem Statement:
//Implement basic math operations  (+, -, *, /, %) on two numbers
//Here, in this problem statement we need perform basic math operations using +,-,*,/,% using Go language
//considering two numbers by using datatypes.Create a seperate concerns for basic math operations.And display
//the reuslt in the main function.

//Algorithm:
//1)Firstly,create package of main and then import fmt
//2)Create seperate concerns for add,sub,mul,div and mod operations
//3)write a logic within that concerns
//4)write a main function by by calling methods including parameters and print the result on terminal

// Code:
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func mul(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}
func mod(a int, b int) int {
	return a % b
}
func main() {
	fmt.Println(add(2, 4))
	fmt.Println(sub(12, 4))
	fmt.Println(mul(2, 4))
	fmt.Println(div(12, 4))
	fmt.Println(mod(12, 4))
}

//Complexity:
//1)The complexity of this program is O(1) Because 1 is constant.
//2)Here, there is no loop.This program consists simple statements.So the complexity is O(n).

//Testcases:
//1)First add 2 numbers if it is positive or negative it will give answer and pass.
//2)By subtracting smallest number to biggest number the result will give negative number and pass.
//3)Multiplying 2 numbers will pass the testcase.
//4)By dividing 2 numbers, while numerator is 12 and denominator is 0 then the result will be error.
//Becasuse no number can be divided by 0.This condition will fail.Dividing any 2 numbers other than zero
//can result pass.
//5)In reminder concept,we have to divide the 2 numbers and we have to write reminder as result and it will pass.
