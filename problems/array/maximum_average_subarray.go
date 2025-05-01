package array

import "math"

/**
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.



Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000

Time Complexity: O(n)
Space Complexity: O(1)
*/

func FindMaxAverage(nums []int, k int) float64 {

	size := len(nums)
	if k > size {
		k = size
	}

	// Initial sum of the first 'k' elements
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	// Set the initial maximum sum as the sum of the first window
	maxSum := sum

	// Use sliding window to update sum and find the maximum sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k] // Add the next element and remove the element that is out of the window
		if sum > maxSum {          // Update max sum if we found a higher sum
			maxSum = sum
		}
	}

	// Return the average of the maximum sum found
	return math.Round(float64(maxSum)/float64(k)*1000) / 1000
}
