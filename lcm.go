package main

import "fmt"

func lcm(m int, n int) int {
	return (m * n) / gcd(m, n)
}
func gcd(m int, n int) int {
	for n != 0 {
		m, n = n, m%n
	}
	return m

}

func main() {
	var m int
	var n int
	fmt.Print("Enter a 1st number:")
	fmt.Scan("%d", &m)
	fmt.Print("Enter a 2nd number:")
	fmt.Scan("%d", &n)
	result := lcm(m, n)
	fmt.Printf("LCM of %d and %d is: %d\n", m, n, result)

}
