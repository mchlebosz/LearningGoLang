package main

import (
	"math"
	"strconv"
)

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func isSuperPrime(x int) bool {
	if !isPrime(x) {
		return false
	}
	digitSum := 0
	for x > 0 {
		digitSum += x % 10
		x /= 10
	}
	return isPrime(digitSum)
}

func toBinary(x int) string {
	Bvalue := ""
	for x > 0 {
		Bvalue = strconv.Itoa(x%2) + Bvalue
		x /= 2
	}
	return Bvalue
}

func isBSuperPrime(x int) bool {
	if !isBSuperPrime(x) {
		return false
	}
	BinaryX := toBinary(x)
	sumX := 0
	for i := 0; i < len(BinaryX); i++ {
		if BinaryX[i] == '1' {
			sumX++
		}
	}
	return isPrime(sumX)
}

func isPerfect(x int) bool {
	sum := 0
	divisor := x - 1
	for divisor > 0 {
		if x%divisor == 0 {
			sum += divisor
		}
		divisor--
	}
	return x == sum
}

func isDivisible(x int) bool {
	sum := 0

	xCopy := x
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return xCopy%sum == 0
}

func isMersenne(x int) bool {
	Bvalue := toBinary(x + 1)
	length := len(Bvalue) - 1
	return isPrime(length) && math.Pow(2, float64(length)) == float64(x+1)
}

func areSiblings(x int, y int) bool {
	return isPrime(x) && isPrime(y) && math.Abs(float64(x-y)) == 2.0
}

func sumDivisors(x int) int {
	sum := 0
	divisor := 1
	for divisor < x {
		if x%divisor == 0 {
			sum += divisor
		}
		divisor++
	}
	return sum
}

func areFriends(x int, y int) bool {
	return x == sumDivisors(y) && y == sumDivisors(x)
}

func euclidesGCD(x int, y int) int {
	for y != 0 {
		z := x % y
		x = y
		y = z
	}
	return x
}

func factorialOf(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorialOf(x-1)
}
