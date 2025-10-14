package utils

import "fmt"

func RomanToInt(s string) int {
	result := 0

	dic := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for _, v := range s {
		fmt.Println(result)
		result += dic[string(v)]
	}

	return result
}
