// @leet start
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for index1, value := range nums {
		needed := target - value
		if index2, exists := seen[needed]; exists {
			return []int{index1, index2}
		} else {
			seen[value] = index1
		}
	}
	return []int{}
}

// @leet end
