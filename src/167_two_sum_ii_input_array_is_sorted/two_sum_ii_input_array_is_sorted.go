package solution

func twoSum(numbers []int, target int) []int {
	if len(numbers) >= 2 {
		x, y := 0, len(numbers)-1

		for x < y {
			if numbers[x]+numbers[y] < target {
				x++
			} else if numbers[x]+numbers[y] > target {
				y--
			} else {
				return []int{x + 1, y + 1}
			}
		}
	}

	return []int{}
}
