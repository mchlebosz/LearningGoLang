package main

import (
	"fmt"
)

func main() {
	value := 5
	fmt.Println("Hello, World!")
	fmt.Println(isEven(value))
	fmt.Println(isPrime(value))
	fmt.Println(isSuperPrime(value))
	fmt.Println(toBinary(value))
	fmt.Println(isSuperPrime(value))
	fmt.Println(isPerfect(value))
	fmt.Println(isDivisible(value))
	fmt.Println(isMersenne(value))
	fmt.Println(areSiblings(value, 13))
	fmt.Println(areFriends(value, 1210))
	fmt.Println(euclidesGCD(value, 96))
	fmt.Println(factorialOf(value))

}
