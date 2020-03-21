package solution

func lengthOfLongestSubstring(s string) int {
	left, ans := 0, 0

	m := map[byte]int{}

	for i := 0; i < len(s); i++ {
		if v, exist := m[s[i]]; exist && left < v {
			left = v
		}

		if ans < (i - left + 1) {
			ans = i - left + 1
		}

		m[s[i]] = i + 1
	}

	return ans
}
