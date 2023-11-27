package main

import "fmt"

func fibonacciseries(n int) {
	var fib1 int = 0
	var fib2 int = 1
	if n == 1 {
		fmt.Print(fib1)
	} else if n == 2 {
		fmt.Print(fib1, fib2)
	} else {
		fmt.Println(fib1, fib2)
		for i := 2; i <= n; i++ {
			var nextno int = fib1 + fib2
			fmt.Println(nextno)
			fib1 = fib2
			fib2 = nextno
		}
	}
}
func main() {
	var n int
	fmt.Println("Enter the number:")
	fmt.Scanf("%d", &n)
	fibonacciseries(n)
}
