package main

import "math"

func main() {
	findMedianSortedArrays([]int{1, 2}, []int{2, 3})
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	halfLen := (m + n + 1) / 2

	// Binary search bounds
	left, right := 0, m

	for left <= right {
		i := (left + right) / 2
		j := halfLen - i

		if i < m && nums1[i] < nums2[j-1] {
			// i is too small, must increase it
			left = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// i is too big, must decrease it
			right = i - 1
		} else {
			// i is perfect

			// Find max of the left side
			var maxOfLeft int
			if i == 0 {
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				maxOfLeft = nums1[i-1]
			} else {
				maxOfLeft = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
			}

			// If the combined length is odd, return max of left side
			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}

			// Find min of the right side
			var minOfRight int
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				minOfRight = int(math.Min(float64(nums1[i]), float64(nums2[j])))
			}

			// If even, return the average of maxOfLeft and minOfRight
			return float64(maxOfLeft+minOfRight) / 2.0
		}
	}

	return 0.0
}
