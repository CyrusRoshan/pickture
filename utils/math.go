package utils

// Euclidean Division Algorithm
func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func SimplifyFraction(numerator, denominator int) (int, int) {
	gcd := GCD(numerator, denominator)
	return numerator / gcd, denominator / gcd
}
