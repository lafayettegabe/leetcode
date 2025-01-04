// @leet start
func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int) // keep the index of each seen char
	max := 0                   // keep lenght of the longest substring
	left := 0                  // keep the index where the current substring starts

	for pos, char := range s {
		if lastPos, exists := seen[char]; exists && lastPos >= left {
			left = lastPos + 1 // move the window after the previous char to keep the values
		}

		seen[char] = pos

		// pos is the end of the substring, left is the start
		if current := pos - left + 1; current > max {
			max = current
		}
	}
	return max
}

// @leet end
