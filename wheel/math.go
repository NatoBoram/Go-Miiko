package wheel

import "math"

// Phi represents the Golden Ratio.
func Phi() float64 {
	return (1 + math.Sqrt(5)) / 2
}

// MaxInt returns the larger of x or y.
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MinInt returns the smaller of x or y.
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// StringInSlice checks if a string is inside a slice.
func StringInSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}
