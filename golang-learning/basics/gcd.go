package main

import "fmt"

func Gcd(a,b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func ShowGcd() {
	a,b := 48,18
	fmt.Printf("GCD of %d and %d is %d\n", a, b, Gcd(a,b))
}
