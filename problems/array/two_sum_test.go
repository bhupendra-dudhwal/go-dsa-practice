package array

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		// Basic test cases
		{"simple match", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"unsorted input", []int{3, 2, 4}, 6, []int{1, 2}},
		{"match at end", []int{1, 2, 3, 4}, 7, []int{2, 3}},

		// Edge cases
		{"empty input", []int{}, 5, []int{}},
		{"no match", []int{1, 2, 3}, 10, []int{}},
		{"single element", []int{5}, 5, []int{}},

		// Advanced/robust cases
		{"duplicates valid pair", []int{3, 3}, 6, []int{0, 1}},
		{"negative numbers", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{"large array", makeLargeTestInput(), 19999, []int{9998, 9999}}, // helper func
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TwoSum(%v,%d)=%v; expected %v", tt.nums, result, tt.target, tt.expected)
			}
		})
	}
}

// Helper function
func makeLargeTestInput() []int {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i + 1
	}
	return nums
}

// Helper function to generate test input
func generateInput(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	return nums
}

// Benchmarking
func BenchmarkTwoSum(b *testing.B) {
	benchmarks := []struct {
		name   string
		nums   []int
		target int
	}{
		{
			name:   "small_input",
			nums:   []int{1, 2, 3, 4, 5},
			target: 9,
		},
		{
			name:   "medium_input",
			nums:   generateInput(1000),
			target: 1999,
		},
		{
			name:   "large_input",
			nums:   generateInput(10000),
			target: 19999,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TwoSum(bm.nums, bm.target)
			}
		})
	}
}
