package utils

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	reversed := 0
	n := x
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed == x
}
