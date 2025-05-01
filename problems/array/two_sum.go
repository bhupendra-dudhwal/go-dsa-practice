package array

/*
*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

Time Complexity: O(n)
Space Complexity: O(n)
*/
func TwoSum(nums []int, target int) []int {
	// Create a map to store numbers and their corresponding indices
	seen := make(map[int]int)

	for currentIndex, val := range nums {
		// Calculate the number needed to reach the target sum
		complement := target - val

		// Check if the complement has already been seen
		if complementIndex, found := seen[complement]; found {
			// If found, return the indices of the complement and current value
			return []int{complementIndex, currentIndex}
		}

		// Store the current number with its index for future lookups
		seen[val] = currentIndex
	}

	// Return an empty slice if no valid pair is found
	return []int{}
}
