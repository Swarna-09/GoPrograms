//Problem Statement:
//Find if a number is prime
//Here, we need to find the given number is prime or not.
//Create seperate concern for finding prime number and write the logic within that method
//and display the output on the terminal.

//Algorithm:
//1)Take the input from the user as n
//2)Check whether the given input from the user is prime or not by using mod operation.
//3)Call the prime method in main method and read the value from the user and display the
//result on terminal.

// Code
package main

import "fmt"

func prime(n int) int {
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			return 0
		}
	}
	return n
}
func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &n)
	if prime(n) == 0 {
		fmt.Println("Prime number")
	}
	fmt.Println("Not a prime number")
}

//Complexity:
//The complexity is O(n).

//Testcases:
//1)If a number is equal to 0 then it is prime number and pass.
//2)Otherwise it is not prime number and fail.
