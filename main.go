package main

import (
	"fmt"
)

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	// Characters representing digits for bases up to 16
	digits := "0123456789ABCDEF"

	// Handle negative values
	isNegative := false
	if value < 0 {
		isNegative = true
		value = -value
	}

	// Convert the value to the specified base
	result := ""
	for value > 0 {
		remainder := value % base
		result = string(digits[remainder]) + result
		value /= base
	}

	// Handle the case where value is 0
	if result == "" {
		result = "0"
	}

	// Add minus sign if the original value was negative
	if isNegative {
		result = "-" + result
	}

	return result
}

func main() {
	fmt.Println(ItoaBase(125, 10))  // "125"
	fmt.Println(ItoaBase(125, 2))   // "1111101"
	fmt.Println(ItoaBase(125, 16))  // "7D"
	fmt.Println(ItoaBase(-125, 10)) // "-125"
	fmt.Println(ItoaBase(-125, 2))  // "-1111101"
	fmt.Println(ItoaBase(-125, 16)) // "-7D"
}
