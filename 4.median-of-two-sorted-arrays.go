// @leet start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArray := make([]int, 0, len(nums1)+len(nums2))

	// Initialize pointers
	var p1, p2 int

	// Add in a sorted way
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			mergedArray = append(mergedArray, nums1[p1])
			p1++
		} else {
			mergedArray = append(mergedArray, nums2[p2])
			p2++
		}
	}

	// Continue adding the remainder array
	for p1 < len(nums1) {
		mergedArray = append(mergedArray, nums1[p1])
		p1++
	}
	for p2 < len(nums2) {
		mergedArray = append(mergedArray, nums2[p2])
		p2++
	}

	n := len(mergedArray)
	if n%2 == 0 {
		// Doesn't have a center, sum both center indices, convert to float and /2
		n1 := mergedArray[n/2-1]
		n2 := mergedArray[n/2]
		return float64(n1+n2) / 2.0
	} else {
		// It has a center
		return float64(mergedArray[n/2])
	}
	return 0.0
}

// @leet end
