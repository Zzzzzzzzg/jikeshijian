func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			for j := i + 1; j < len(digits); j++ {
				digits[j] = 0
			}
		return digits
		}
	}
	digits = make([]int, len(digits) + 1)
	digits[0] = 1
	return digits
}
