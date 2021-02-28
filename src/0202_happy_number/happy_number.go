package solution

func isHappy(n int) bool {
	m := make(map[int]bool)

	sum := 0

	for n != 1 {
		for n > 0 {
			x := n % 10
			n /= 10
			sum += x * x
		}

		n = sum
		sum = 0

		if _, exist := m[n]; exist {
			return false
		}

		m[n] = true
	}

	return true
}
